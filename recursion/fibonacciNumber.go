package recursion

var cache = map[int]int{}
//The Fibonacci numbers, commonly denoted F(n) form a sequence,
// called the Fibonacci sequence, such that each number is the sum of the two preceding ones,
// starting from 0 and 1. That is,
//
//F(0) = 0, F(1) = 1
//F(n) = F(n - 1) + F(n - 2), for n > 1.
//Given n, calculate F(n).
// https://leetcode.com/problems/fibonacci-number/
// recursion + memorization
func fib(n int) int {
	val, ok := cache[n]
	if ok {
		return val
	}
	if n < 2 {
		return n
	}
	res := fib(n-1) + fib(n-2)
	cache[n] = res
	return res
}
