package slidingwindow

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func LengthOfLongestSubstring(s string) int {
	// to char ints
	runes := []rune(s)
	charCnt := map[rune]int{} // letter count map
	rep := 0 // current number of repeated letter
	l := 0 // left position of the window
	max := 0 // max length
	for i, v := range runes {
		cnt, _ := charCnt[v]
		charCnt[v] = cnt + 1
		// found repeated letter
		if cnt > 0 {
			rep++
		}
		// move left position until no repeated letter
		for ; rep > 0; {
			cnt, _ := charCnt[runes[l]]
			charCnt[runes[l]]--
			if cnt == 2 {
				rep--
			}
			l++
		}
		// compare length once no repeated letter
		if i - l + 1 > max {
			max = i - l + 1
		}
	}
	return max
}
