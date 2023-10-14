package tree_search

import (
	"fmt"
	"math"
	"task/collections"
)

func DijkstraRelax(startNode int, goalNode int, graph map[int][]collections.PathTransition) []int {
	// dist := make([]int, len(graph))
	used := make([]uint, len(graph))
	prev := make([]int, len(graph))
	// var path []int

	priorityQueue := collections.NewPriorityQueue()
	nodeMap := make(map[int]*collections.QueueNode)

	// for i := range graph {
	// 	if i != startNode {
	// 		priorityQueue.Push(&collections.QueueNode{i, math.MaxInt})
	// 	}
	// }
	priorityQueue.Push(&collections.QueueNode{Station: startNode, Time: 0})
	nodeMap[startNode] = &collections.QueueNode{Station: startNode, Time: 0}

	for i := 0; i < len(graph); i++ {
		if i != startNode {
			nodeMap[i] = &collections.QueueNode{Station: startNode, Time: math.MaxInt}
		}
	}
	prev[startNode] = -1
	// dist[startNode] = 0

	for priorityQueue.Size() > 0 {
		node, _ := priorityQueue.Pop()

		vertex := node.(*collections.QueueNode).Station
		distance := node.(*collections.QueueNode).Time

		used[vertex] = 1

		// var usedLocal int
		var minTransitionCost int

		for _, transition := range graph[vertex] {
			minTransitionCost = transition.Time
			for _, transitionLocal := range graph[vertex] {
				if transition.SecondPoint == transitionLocal.SecondPoint && minTransitionCost > transition.Time {
					minTransitionCost = transition.Time
				}
			}
			// usedLocal = transition.SecondPoint

			fmt.Println(minTransitionCost)

			if used[transition.SecondPoint] == 0 {
				uNode, nodeIsPresent := nodeMap[transition.SecondPoint]
				if !nodeIsPresent {
					nodeMap[transition.SecondPoint] = &collections.QueueNode{
						Station: transition.SecondPoint,
						Time:    distance + minTransitionCost,
					}
				} else {
					if uNode.Time > distance+minTransitionCost {
						uNode.Station = transition.SecondPoint
						uNode.Time = distance + minTransitionCost
					}
				}
				priorityQueue.Push(
					&collections.QueueNode{
						Station: transition.SecondPoint,
						Time:    distance + minTransitionCost,
					},
				)
			}
		}
	}

	result := make([]int, len(nodeMap))

	for i, node := range nodeMap {
		result[i] = node.Time
	}

	return result
	// return -1, make([]int, 0)
}
