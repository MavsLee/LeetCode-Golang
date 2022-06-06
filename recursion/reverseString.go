package recursion

// https://leetcode.com/problems/reverse-string/
func reverseString(s []byte)  {
	helper(0, len(s) - 1, s)
}

func helper(start int, end int, s []byte) {
	if start >= end {
		return
	}
	// swap start and end
	temp := s[start]
	s[start] = s[end]
	s[end] = temp
	helper(start+1, end - 1, s)
}
