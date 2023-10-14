package tree_search

import (
	"task/collections"
)

func BreadthFirstSearch(startNode int, goalNode int, graph map[int][]int) (int, []int) {
	used := make([]uint, len(graph))
	prev := make([]int, len(graph))
	var path []int

	queue := collections.NewQueue[collections.QueueNode]()

	queue.AddLast(collections.QueueNode{Station: startNode, Time: 0})
	prev[startNode] = -1

	for !queue.IsEmpty() {
		node, _ := queue.PollFirst()

		if node.Station == goalNode {

			current := goalNode
			path = append(path, current)

			for prev[current] != -1 {
				current = prev[current]
				path = append(path, current)
			}

			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}

			return node.Time, path
		}

		used[node.Station] = 1

		for _, child := range graph[node.Station] {
			if used[child] == 0 {
				queue.AddLast(collections.QueueNode{Station: child, Time: node.Time + 1})
				used[child] = 1
				prev[child] = node.Station
			}
		}

	}

	return -1, make([]int, 0)
}
