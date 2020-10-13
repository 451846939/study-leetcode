package main

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func swapPairs(head *ListNode) *ListNode {
	return reverseKGroup(head, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	tmp := &ListNode{Next: head}
	pre := tmp
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return tmp.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return tmp.Next
}

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	tmp := head
	for prev != tail {
		next := tmp.Next
		tmp.Next = prev
		prev = tmp
		tmp = next
	}

	return tail, head
}
