package main

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/balanced-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := dfs(root.Left, 0)
	r := dfs(root.Right, 0)
	return abs(r, l) && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(r int, l int) bool {
	if r-l > 1 || l-r > 1 {
		return false
	}
	return true
}

func dfs(node *TreeNode, count int) int {
	if node == nil {
		return count
	}
	r := dfs(node.Right, count+1)
	l := dfs(node.Left, count+1)
	return max(r, l)
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
