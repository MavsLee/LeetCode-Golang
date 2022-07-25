package graph

func findCircleNum(isConnected [][]int) int {
	// union find
	// find parent save to map
	length := len(isConnected)
	if length == 0 {
		return 0
	}
	uf := NewUF(length)
	for i, v := range isConnected {
		for j, k := range v {
			if k == 0 || i == j {
				continue
			}
			uf.Connect(i, j)
		}
	}
	res := map[int]bool{}
	for i := 0; i < length; i++ {
		p := uf.Find(i)
		_, ok := res[p]
		if !ok {
			res[p] = true
		}
	}
	return len(res)
}
