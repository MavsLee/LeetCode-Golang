package slidingwindow

import "math"

// for a specific position, how much water it can hold depends on min(maxLeft, maxRight) - height
// https://leetcode.com/problems/trapping-rain-water/
func trap(height []int) int {
	maxLH := make([]int, len(height))
	maxl := 0
	for i, v := range height {
		maxLH[i] = maxl
		if v > maxl {
			maxl = v
		}
	}
	res := 0
	maxR := 0
	for i:=len(height) - 1; i>=0; i-- {
		temp := int(math.Min(float64(maxLH[i]), float64(maxR))) - height[i]
		if temp > 0 {
			res += temp
		}
		if height[i] > maxR {
			maxR = height[i]
		}
	}
	return res
}
