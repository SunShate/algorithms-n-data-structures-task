package main

import (
	"bytes"
	"errors"
	"fmt"
	"sort"

	"task/collections"
)

type trajectory []int

/**
 * Implement scanner for wrapping input into integer array
 * Try: 1 3 2 5 3 4
 */
func (s *trajectory) Scan(state fmt.ScanState, verb rune) error {
	tok, err := state.Token(false, func(r rune) bool { return r != '\n' })
	if err != nil {
		return err
	}
	if _, _, err := state.ReadRune(); err != nil {
		if len(tok) == 0 {
			panic(err)
		}
	}
	b := bytes.NewReader(tok)
	for {
		var d int
		_, err := fmt.Fscan(b, &d)
		if err != nil {
			break
		}
		*s = append(*s, d)
	}
	return nil
}

type Neighbour struct {
	Node  int   // neighbour node
	Buses []int // buses that are going to this stop
}

var routs map[int][]int
var path []int

func findNeighbour(neighbours []Neighbour, targetNode int) *Neighbour {
	for i := range neighbours {
		if neighbours[i].Node == targetNode {
			return &neighbours[i]
		}
	}
	return nil
}

func dfs(graph [][]Neighbour, nodeCur, nodeLast int) {
	if !contains(path, nodeCur) {
		path = append(path, nodeCur)

		if nodeCur == nodeLast {
			routs[len(routs)] = append([]int{}, path...)
			path = path[:len(path)-1]
			return
		}

		for _, val := range graph[nodeCur-1] {
			dfs(graph, val.Node, nodeLast)
		}
		path = path[:len(path)-1]
	}
}

func contains(slice []int, element int) bool {
	for _, val := range slice {
		if val == element {
			return true
		}
	}
	return false
}

func main() {
	var stations, trajectories, startStation, endStation int

	fmt.Print("Write amount of stations: ")
	fmt.Scanln(&stations)
	fmt.Print("Write amount of trajectories: ")
	fmt.Scanln(&trajectories)

	stationsSet := make(collections.Set)
	var busRoutes []trajectory

	for i := 0; i < trajectories; i++ {
		var Trajectory trajectory
		fmt.Printf("Write trajectory %d: ", i+1)
		fmt.Scan(&Trajectory)
		busRoutes = append(busRoutes, Trajectory)

		stationsSet.AddMulti(Trajectory...)
	}

	if stationsSet.Size() > stations {
		fmt.Println(errors.New("Trajectory has more stations than determined"))
	}

	fmt.Print("Write start station: ")
	fmt.Scanln(&startStation)

	if !stationsSet.Has(startStation) {
		fmt.Println(errors.New("This start station not exist"))
	}

	fmt.Print("Write end station: ")
	fmt.Scanln(&endStation)

	if !stationsSet.Has(endStation) {
		fmt.Println(errors.New("This end station not exist"))
	}

	graph := make([][]Neighbour, 0)
	for vert := 0; len(graph) == vert; {
		vert++
		for r, route := range busRoutes {
			for c, stop := range route {
				if stop == vert {
					if len(graph) < vert {
						graph = append(graph, []Neighbour{})
					}

					if c > 0 { // if not first then search for left neighbour
						leftNeighbour := findNeighbour(graph[vert-1], route[c-1])
						if leftNeighbour != nil { // if neighbour from this bus route found then add it's upcomming busses
							leftNeighbour.Buses = append(leftNeighbour.Buses, r)
						} else { // if neighbour from this bus route not found
							graph[vert-1] = append(graph[vert-1], Neighbour{
								Node:  route[c-1],
								Buses: []int{r},
							})
						}
					}
					if c < len(route)-1 { // if not last then search for right neighbour
						rightNeighbour := findNeighbour(graph[vert-1], route[c+1])
						if rightNeighbour != nil {
							rightNeighbour.Buses = append(rightNeighbour.Buses, r)
						} else {
							graph[vert-1] = append(graph[vert-1], Neighbour{
								Node:  route[c+1],
								Buses: []int{r},
							})
						}
					}
				}
			}
		}
	}

	for i, vertexes := range graph {
		fmt.Println("Neighbour stops for ", i, ":", vertexes)
	}
	routs = make(map[int][]int)
	dfs(graph, startStation, endStation)

	fmt.Println("\nRoutes:")
	for _, route := range routs {
		fmt.Println(route)
	}

	paths := make([][][]int, len(routs))
	for j, route := range routs {
		for i := 0; i < len(route)-1; i++ {
			for _, val := range graph[route[i]-1] {
				if val.Node == route[i+1] { // if next node in neighbour stops for (route[i] - 1) is node from route then append node incomming busses
					paths[j] = append(paths[j], val.Buses)
					break
				}
			}
		}
	}

	fmt.Println("\nBuses:")
	for _, path := range paths {
		for _, nodes := range path {
			fmt.Println(nodes)
		}
		fmt.Println()
	}

	maxTrajLen := make([][]struct {
		number, length int
	}, len(paths))
	for j, path := range paths {
		i := 0
		for i < len(path) {
			maxLen := 1
			var trajLen []struct {
				number, length int
			}
			for _, busNumber := range path[i] {
				length := 0
				for k := i; k < len(path); k++ {
					if contains(path[k], busNumber) {
						length++
					} else {
						break
					}
				}
				trajLen = append(trajLen, struct {
					number, length int
				}{busNumber, length})
			}
			sort.Slice(trajLen, func(i, j int) bool {
				return trajLen[i].length < trajLen[j].length
			})
			maxLen = trajLen[len(trajLen)-1].length
			for _, pr := range trajLen {
				if pr.length == maxLen {
					maxTrajLen[j] = append(maxTrajLen[j], pr)
					break
				}
			}
			i += maxLen
		}
	}

	for i, rout := range maxTrajLen {
		fmt.Printf("Route № %d\n", i)
		for _, pr := range rout {
			fmt.Printf("TrajNum: %d \t Length: %d\n", pr.number, pr.length)
		}
	}

	weights := make([]int, len(maxTrajLen))
	for i, rout := range maxTrajLen {
		ribs := 0
		transfers := len(rout) - 1
		for _, pr := range rout {
			ribs += pr.length
		}
		weight := ribs + 3*transfers
		weights[i] = weight
		fmt.Printf("Route № %d Time: %d\n", i, weight)
	}

	ind := 0
	for i, w := range weights {
		if w < weights[ind] {
			ind = i
		}
	}

	fmt.Println("\nAnswer:")
	for _, pr := range maxTrajLen[ind] {
		fmt.Printf("Bus number: %d \t Stops count: %d\n", pr.number, pr.length)
	}
	fmt.Println()
}
