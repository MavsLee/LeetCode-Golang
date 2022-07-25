package graph

import (
	"container/heap"
	"fmt"
	"math"
)

func minCostConnectPoints(points [][]int) int {
	// get all edges, calcuate weight by mahhtan distance
	//krushal algori, sort by edges weight, choose smallest one if not forming cycyle, until N-1 is added
	// user indice as vertex label
	if len(points) == 0 {
		return 0
	}
	size := len(points)
	uf := NewUF(size)

	var edges PriorityQueue
	for i, _ := range points {
		for j := i + 1; j < len(points); j++ {
			dis := int(math.Abs(float64(points[i][0]-points[j][0])) + math.Abs(float64(points[i][1]-points[j][1])))
			edge := &Edge{
				V1:   i,
				V2:   j,
				Cost: dis,
			}
			edges = append(edges, edge)
			//fmt.Println("create edge", edge.V1, edge.V2, edge.Cost)
		}
	}
	//fmt.Println("total edges", len(edges))
	//pq := make(PriorityQueue, len(edges))
	//for _, v := range edges {
	//	pq = append(pq, v)
	//}
	//fmt.Println("total pq", len(pq))
	heap.Init(&edges)
	res := 0
	count := size - 1 // N-1 edges
	for len(edges) > 0 {
		cur := heap.Pop(&edges).(*Edge)
		//fmt.Println("pop edge", cur.V1, cur.V2, cur.Cost)
		if !uf.IsConnected(cur.V1, cur.V2) {
			uf.Connect(cur.V1, cur.V2)
			count--
			res += cur.Cost
			//fmt.Println("res", res)
		}
	}
	return res
}

type Edge struct {
	V1   int
	V2   int
	Cost int
}

type PriorityQueue []*Edge

func (p *PriorityQueue) Len() int {
	return len(*p)
}

// min-heap
func (p *PriorityQueue) Less(i, j int) bool {
	//fmt.Println("pq less", i, j)
	if p == nil {
		fmt.Println("pq is nil")
	}
	if (*p)[i] == nil {
		fmt.Println("vertex i is nil", (*p)[i], i)
	}
	if (*p)[j] == nil {
		fmt.Println("vertex j is nil", (*p)[j], j)
	}
	return (*p)[i].Cost < (*p)[j].Cost
}

func (p *PriorityQueue) Swap(i, j int) {
	temp := (*p)[i]
	(*p)[i] = (*p)[j]
	(*p)[j] = temp
	//p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	item := x.(*Edge)
	//item.index = n
	*p = append((*p), item)
}

func (p *PriorityQueue) Pop() interface{} {
	//old := p
	//n := len(old)
	//item := old[n-1]
	//old[n-1] = nil // avoid memory leak
	//item.index = -1 // for safety
	//p = old[0 : n-1]

	n := p.Len()
	res := (*p)[n-1]
	*p = (*p)[:n-1]
	return res
}
