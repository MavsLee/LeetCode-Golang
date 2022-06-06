package recursion

// https://leetcode.com/problems/pascals-triangle-ii/
// recursion
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	preVals := getRow(rowIndex - 1)
	var res []int
	for i, v := range preVals {
		if i == 0 {
			res = append(res, 1)
		}
		if i == len(preVals) - 1 {
			res = append(res, 1)
			return res
		}
		res = append(res, v + preVals[i+1])
	}
	return res
}
