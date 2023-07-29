package cv

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var cnt int
	var prefixFind func(node *TreeNode)
	prefixFind = func(node *TreeNode) {
		dfs(node, targetSum, &cnt)
		prefixFind(node.Left)
		prefixFind(node.Right)
	}
	prefixFind(root)

	return cnt
}

func dfs(root *TreeNode, targetSum int, cnt *int) {
	if root == nil {
		return
	}

	if root != nil && root.Val == targetSum {
		*cnt = *cnt + 1
	}

	dfs(root.Left, targetSum-root.Val, cnt)
	dfs(root.Right, targetSum-root.Val, cnt)
}

func findMode(root *TreeNode) (res []int) {

	var base, cnt, maxCnt int

	var update func(val int)

	update = func(val int) {
		if val != base {
			base = val
			cnt = 1
		} else {
			cnt++
			if cnt == maxCnt {
				res = append(res, []int{val}...)
			} else if cnt > maxCnt {
				res = []int{val}
				maxCnt = cnt
			}
		}
	}

	var inorder func(*TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		update(node.Val)
		inorder(node.Right)

	}
	return
}
