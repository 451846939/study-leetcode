package main

/*

给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
*/
func main() {

}

/**
 * Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	l := 0
	for h := head; h != nil; h = h.Next {
		l++
	}
	return toBST(0, l-1, &head)
}

func toBST(l int, r int, nodePrt **ListNode) *TreeNode {
	if l > r {
		return nil
	}
	//中序遍历
	mid := l + ((r - l) >> 1)
	lTree := toBST(l, mid-1, nodePrt)
	val := (*nodePrt).Val
	*nodePrt = (*nodePrt).Next
	return &TreeNode{
		Val:   val,
		Left:  lTree,
		Right: toBST(mid+1, r, nodePrt),
	}
}
