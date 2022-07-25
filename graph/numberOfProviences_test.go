package graph

import (
	"fmt"
	"testing"
)

func Test_findCircleNum(t *testing.T) {
	input := [][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1},
	}
	res := findCircleNum(input)
	fmt.Println(res)
}
