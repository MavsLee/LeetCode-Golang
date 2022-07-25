package graph

func countComponents(n int, edges [][]int) int {
	uf := NewUF(n)
	for _, e := range edges {
		uf.Connect(e[0], e[1])
	}
	res := uf.GetConnectedComponents()
	return res
}
