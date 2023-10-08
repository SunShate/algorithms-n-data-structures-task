package main

import (
	"bytes"
	"errors"
	"fmt"

	"task/collections"
)

type PathTransition struct {
	firstPoint       int
	secondPoint      int
	time             int
	isBackwards      bool
	trajectoryNumber int
	isTransfer       bool
}

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

	var graph []PathTransition
	used := make([]bool, stations)

	// fill the graph with transition between stations
	for trajNum, traj := range busTrajectories {
		for i := 0; i < len(traj)-1; i++ {
			if used[traj[i]] {
				for _, pathTransition := range graph {
					if traj[i] == pathTransition.secondPoint && trajNum != pathTransition.trajectoryNumber && !pathTransition.isTransfer {
						if pathTransition.isBackwards {
							graph = append(graph, PathTransition{pathTransition.secondPoint, pathTransition.firstPoint, 3, !pathTransition.isBackwards, trajNum, true})
							graph = append(graph, PathTransition{pathTransition.firstPoint, pathTransition.secondPoint, 3, pathTransition.isBackwards, trajNum, true})
						} else {
							graph = append(graph, PathTransition{pathTransition.firstPoint, pathTransition.secondPoint, 3, pathTransition.isBackwards, trajNum, true})
							graph = append(graph, PathTransition{pathTransition.secondPoint, pathTransition.firstPoint, 3, !pathTransition.isBackwards, trajNum, true})
						}
					}
				}
			}
			graph = append(graph, PathTransition{firstPoint: traj[i], secondPoint: traj[i+1], time: 1, isBackwards: false, trajectoryNumber: trajNum})
			graph = append(graph, PathTransition{firstPoint: traj[i+1], secondPoint: traj[i], time: 1, isBackwards: true, trajectoryNumber: trajNum})

			used[traj[i]] = true
			used[traj[i+1]] = true

			if i == len(traj)-2 {
				fmt.Println("s")
			}

			if used[traj[i+1]] && i == len(traj)-2 {
				for _, pathTransition := range graph {
					if traj[i+1] == pathTransition.secondPoint && trajNum != pathTransition.trajectoryNumber && !pathTransition.isTransfer {
						if pathTransition.isBackwards {
							graph = append(graph, PathTransition{pathTransition.secondPoint, pathTransition.firstPoint, 3, !pathTransition.isBackwards, trajNum, true})
							graph = append(graph, PathTransition{pathTransition.firstPoint, pathTransition.secondPoint, 3, pathTransition.isBackwards, trajNum, true})
						} else {
							graph = append(graph, PathTransition{pathTransition.firstPoint, pathTransition.secondPoint, 3, pathTransition.isBackwards, trajNum, true})
							graph = append(graph, PathTransition{pathTransition.secondPoint, pathTransition.firstPoint, 3, !pathTransition.isBackwards, trajNum, true})
						}
					}
				}
			}
		}
	}

	for _, transition := range graph {
		if transition.isTransfer {
			fmt.Println(transition)
		}
	}
	// cost, path := tree_search.BreadthFirstSearch(startStation, endStation, graph)
	// fmt.Println("Time to get to the end station:", cost)
	// fmt.Println("Path:", path)
}
