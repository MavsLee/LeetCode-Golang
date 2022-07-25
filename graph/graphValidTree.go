package graph

func validTree(n int, edges [][]int) bool {
	if n <= 1 {
		return true
	}
	if len(edges) != n-1 {
		return false
	}
	uf := NewUF(n)
	for _, e := range edges {
		uf.Connect(e[0], e[1])
	}
	p := uf.Find(0)
	for i := 1; i < n; i++ {
		ip := uf.Find(i)
		if ip != p {
			return false
		}
	}
	return true
}
