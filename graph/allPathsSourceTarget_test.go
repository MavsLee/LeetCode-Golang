package graph

import (
	"fmt"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	input := [][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}
	res := allPathsSourceTarget(input)
	fmt.Println(res)
}
