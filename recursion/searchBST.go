package recursion

// https://leetcode.com/problems/search-in-a-binary-search-tree/
// https://leetcode.com/problems/pascals-triangle-ii/
func searchBST(root *BinaryTreeNode, val int) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
