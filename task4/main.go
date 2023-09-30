package main

import (
	"bytes"
	"errors"
	"fmt"

	"task/collections"
	"task/tree_search"
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

func main() {
	var stations, trajectories, startStation, endStation int

	fmt.Print("Write amount of stations: ")
	fmt.Scanln(&stations)
	fmt.Print("Write amount of trajectories: ")
	fmt.Scanln(&trajectories)

	stationsSet := make(collections.Set)
	var busTrajectories []trajectory

	for i := 0; i < trajectories; i++ {
		var Trajectory trajectory
		fmt.Printf("Write trajectory %d: ", i+1)
		fmt.Scan(&Trajectory)
		busTrajectories = append(busTrajectories, Trajectory)

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

	graph := make(map[int][]int)

	// fill the graph with trajectories
	for _, traj := range busTrajectories {
		for i := 0; i < len(traj)-1; i++ {
			graph[traj[i]] = append(graph[traj[i]], traj[i+1])
			graph[traj[i+1]] = append(graph[traj[i+1]], traj[i])
		}
	}

	fmt.Println(graph)
	fmt.Println("Time to get to the end station", tree_search.BreadthFirstSearch(startStation, endStation, graph))
}
