package slidingwindow

import "math"

//https://leetcode.com/problems/container-with-most-water/
// two pointer, Greedy
// The widest container (using first and last line) is a good candidate, because of its width. Its water level is the height of the smaller one of first and last line.
//All other containers are less wide and thus would need a higher water level in order to hold more water.
//The smaller one of first and last line doesn't support a higher water level and can thus be safely removed from further consideration.
func maxArea(height []int) int {
	len := len(height)
	if len < 2 {
		return 0
	}
	res := 0
	l := 0
	r := len - 1
	for ; l < r; {
		area := float64(r - l) * math.Min(float64(height[l]), float64(height[r]))
		if int(area) > res {
			res = int(area)
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
