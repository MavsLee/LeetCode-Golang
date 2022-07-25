package graph

import "sort"

// // 2-7-5-0-3-4-1
func earliestAcq(logs [][]int, n int) int {
	uf := NewUF(n)
	res := -1
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})
	for _, l := range logs {
		uf.Connect(l[1], l[2])

		res = l[0]
		if uf.GetConnectedComponents() == 1 {
			return res
		}
	}
	if uf.GetConnectedComponents() != 1 {
		return -1
	}
	return res
}
