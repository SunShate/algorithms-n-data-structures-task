package tree_search

import "task/collections"

func BreadthFirstSearch(startNode int, goalNode int, graph map[int][]int) int {
	visitied := make(map[int]bool, len(graph))

	queue := collections.NewQueue[collections.QueueNode]()

	queue.AddLast(collections.QueueNode{Station: startNode, Time: 0})

	for !queue.IsEmpty() {
		node, _ := queue.PollFirst()

		if node.Station == goalNode {
			return node.Time
		}

		visitied[node.Station] = true

		for _, child := range graph[node.Station] {
			if visitied[child] == false {
				queue.AddLast(collections.QueueNode{Station: child, Time: node.Time + 1})
				visitied[child] = true
			}
		}
	}

	return -1
}
