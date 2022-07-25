package graph

import (
	"fmt"
	"testing"
)

func Test_minCostConnectPoints(t *testing.T) {
	input := [][]int{
		{0, 0},
		{2, 2},
		{3, 10},
		{5, 2},
		{7, 0},
	}
	res := minCostConnectPoints(input)
	fmt.Println(res)
}
