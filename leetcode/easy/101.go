package easy

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func isSymmetric(root *TreeNode2) bool {
	if root == nil {
		return true
	}

	var isMirror func(*TreeNode2, *TreeNode2) bool

	isMirror = func(left, right *TreeNode2) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		return isMirror(left.Left, right.Right) &&
			isMirror(left.Right, right.Left)
	}

	return isMirror(root.Left, root.Right)
}
