package recursion

// You are climbing a staircase. It takes n steps to reach the top.
//
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
//
//
//Example 1:
//
//Input: n = 2
//Output: 2
//Explanation: There are two ways to climb to the top.
//1. 1 step + 1 step
//2. 2 steps
//Example 2:
//
//Input: n = 3
//Output: 3
//Explanation: There are three ways to climb to the top.
//1. 1 step + 1 step + 1 step
//2. 1 step + 2 steps
//3. 2 steps + 1 step

// https://leetcode.com/problems/climbing-stairs/
// f(n) = f(n-1) + f(n-2)
// f(1) = 1, f(2) = 2, f(3) = f(1) + f(2)

//var cache = map[int]int{}

func climbStairs(n int) int {
	val, ok := cache[n]
	if ok {
		return val
	}
	if n <= 2 {
		return n
	}
	res := climbStairs(n-1) + climbStairs(n-2)
	cache[n] = res
	return res
}