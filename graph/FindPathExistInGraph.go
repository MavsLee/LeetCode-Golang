package graph

func validPath(n int, edges [][]int, source int, destination int) bool {
	// construct a graph
	// map[int]map[int]int
	// DFS to start from source, check if can reach to dest
	// have visit to mark visited node
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	visited := make([]int, n)
	dfs(graph, visited, source, destination)
	if visited[destination] == 1 {
		return true
	}
	return false
}

// dfs recursion
func dfs(graph [][]int, visited []int, cur int, dest int) {
	if cur == dest {
		return
	}
	visited[cur] = 1
	for _, n := range graph[cur] {
		if visited[n] != 1 {
			dfs(graph, visited, n, dest)
		}
	}
}

// dfs stack, non-recursive
func validPath2(n int, edges [][]int, source int, destination int) bool {
	// construct a graph
	// map[int]map[int]int
	// DFS to start from source, check if can reach to dest
	// have visit to mark visited node
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	visited := make([]bool, n)
	var stack []int
	stack = append(stack, source)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur == destination {
			return true
		}
		if visited[cur] {
			continue
		}
		visited[cur] = true
		for _, n := range graph[cur] {
			stack = append(stack, n)
		}
	}
	return false
}

// BFS with queue
func validPath_BFS(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	var queue []int
	queue = append(queue, source)
	visited := make([]bool, n)
	visited[source] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == destination {
			return true
		}
		for _, n := range graph[cur] {
			if !visited[n] {
				queue = append(queue, n)
				visited[n] = true
			}
		}
	}
	return false
}
