package main

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤m≤n≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dumNode := &ListNode{}
	dumNode.Next = head
	pre := dumNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	leftNode := pre.Next
	cur := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil

	reverseLinkedList := func(head *ListNode) {
		//pre ->cur ->next
		var pre *ListNode
		cur := head
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
	}
	reverseLinkedList(leftNode)

	pre.Next = rightNode
	leftNode.Next = cur

	return dumNode.Next
}
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	dum := &ListNode{}
	dum.Next = head
	pre := dum
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	//		--next----
	//		|		|
	//		v		|
	// pre->cur ->next->next.next
	//		|			    ^
	//		-----next-------|
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		//滚动，所以不能等于cur 需要每次都等于pre的下一个节点next
		next.Next = pre.Next
		pre.Next = next
	}

	return dum.Next
}
