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

func largestValues(root *TreeNode) []int {

	var preorder func(*TreeNode, int)

	levelMap := map[int]int{}
	preorder = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if v, ok := levelMap[level]; !ok {
			levelMap[level] = node.Val
		} else if node.Val > v {
			levelMap[level] = v
		}

		preorder(node.Left, level+1)

		preorder(node.Right, level+1)
	}

	res := make([]int, len(levelMap))

	for i := range res {
		res[i] = levelMap[i]
	}
	return res

}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == root2 {
		return true
	}
	if root1 != nil && root2 == nil || root2 != nil && root1 == nil || root1.Val != root2.Val {
		return false
	}

	return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) || flipEquiv(root1.Right, root2.Left) && flipEquiv(root1.Left, root2.Right)
}
