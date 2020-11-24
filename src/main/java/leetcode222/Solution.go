package main

import (
	"fmt"
)

/*
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-complete-tree-nodes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(countNodes(&TreeNode{1, &TreeNode{2, nil, nil}, nil}))
}

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes1(root *TreeNode) int {
	count := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		count++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return count
}
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	var exists func(node *TreeNode, mid int) bool
	exists = func(node *TreeNode, mid int) bool {
		bits := 1 << (level - 1)
		for node != nil && bits > 0 {
			if bits&mid == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node != nil
	}
	low, high := 1<<level, (1<<(level+1))-1
	for low < high {
		mid := (high-low+1)>>1 + low
		if exists(root, mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}
