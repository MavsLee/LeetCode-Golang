package graph

import "fmt"

var res [][]int
var path []int

func allPathsSourceTarget(graph [][]int) [][]int {
	end := len(graph) - 1
	path = append(path, 0)
	dfs1(graph, 0, end)
	fmt.Println("result:", res)
	return res
}

func dfs1(graph [][]int, cur int, end int) {
	if cur == end {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		//res = append(res, path)

		fmt.Println("find route", path, "result", res)
		return
	}
	for _, n := range graph[cur] {
		path = append(path, n)
		fmt.Println("add to path", n, "result", res, "path", path)
		dfs1(graph, n, end)
		path = path[:len(path)-1]
		fmt.Println("remove from path", n, "result", res, "path", path)
	}
}

func allPathsSourceTargetBFD(graph [][]int) [][]int {
	end := len(graph) - 1
	var rest [][]int
	var queue [][]int
	queue = append(queue, []int{0})
	for len(queue) > 0 {
		cur := queue[0]
		lastNode := cur[len(cur)-1]
		// check cur and remove cur from the queue
		queue = queue[1:]
		if lastNode == end {
			// find a path
			p := make([]int, len(cur))
			copy(p, cur)
			rest = append(rest, p)
			continue
		}
		// visit its neighbours
		for _, n := range graph[lastNode] {
			newCur := make([]int, len(cur))
			copy(newCur, cur)
			newCur = append(newCur, n)
			queue = append(queue, newCur)
		}
	}
	return rest
}
