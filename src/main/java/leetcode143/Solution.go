package main

/*
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func main() {

}

/**
 * Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	midNode := findMidNode(head)
	rNode := midNode.Next
	rNode = reverseList(rNode)
	midNode.Next = nil
	mergeList(head, rNode)
}
func mergeList(n1, n2 *ListNode) {
	var n1Temp, n2Temp *ListNode
	for n1 != nil && n2 != nil {
		n1Temp = n1.Next
		n2Temp = n2.Next

		n1.Next = n2
		n1 = n1Temp

		n2.Next = n1
		n2 = n2Temp
	}
}
func findMidNode(head *ListNode) *ListNode {
	slow, quick := head, head
	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	// 1 cur-> 2  next -> 3
	var prev, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
