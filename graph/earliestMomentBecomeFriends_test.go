package graph

import (
	"fmt"
	"testing"
)

func Test_earliestAcq(t *testing.T) {

	input := [][]int{
		{9, 0, 3},
		{0, 2, 7},
		{12, 3, 1},
		{5, 5, 2},
		{3, 4, 5},
		{1, 5, 0},
		{6, 2, 4},
		{2, 5, 3},
		{7, 7, 3},
	}
	res := earliestAcq(input, 8)
	fmt.Println(res)
}
