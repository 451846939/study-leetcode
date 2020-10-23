package main

import "fmt"

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(isPalindrome(&ListNode{Val: 1, Next: &ListNode{0, &ListNode{Val: 1}}}))
}

/**
 * Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	midNode := findMidNode(head)
	rNode := midNode.Next
	midNode.Next = nil
	reverseRNode := reverseList(rNode)
	l, r := head, reverseRNode
	for l != nil && r != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}
	return true
}

func findMidNode(head *ListNode) *ListNode {
	s, f := head, head
	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	return s
}

func reverseList(head *ListNode) *ListNode {
	//prev -> cur ->next
	var prev, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
