package main

/*
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, res := forTree(root)
	return res
}

func forTree(nowTree *TreeNode) (int, int) {
	if nowTree == nil {
		return 0, 0
	}
	l := 0
	r := 0
	lsum := 0
	rsum := 0
	if nowTree.Left != nil {
		l, lsum = forTree(nowTree.Left)
	}
	if nowTree.Right != nil {
		r, rsum = forTree(nowTree.Right)
	}
	return max(l+1, r+1), max(max(lsum, rsum), l+r)
}

func max(i int, j int) int {
	if i < j {
		return j
	}
	return i
}
