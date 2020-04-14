package main

import (
	"container/list"
	"fmt"
	"strconv"
)

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。



进阶：

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。



示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	//l:=[]int{2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9}
	l := []int{5}
	//l2:=[]int{5,6,4,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9,9,9,9}
	l2 := []int{5}
	root := &ListNode{}
	tmp := root
	for i := range l {
		tmp.Next = &ListNode{l[i], nil}
		tmp = tmp.Next
	}
	root2 := &ListNode{}
	tmp2 := root2
	for i := range l2 {
		tmp2.Next = &ListNode{l2[i], nil}
		tmp2 = tmp2.Next
	}
	numbers := addTwoNumbers(root.Next, root2.Next)
	fmt.Println(numbers)
}

/**
 * Definition for singly-linked list.*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := list.New()
	num2 := list.New()
	for ; l1 != nil; l1 = l1.Next {
		num1.PushBack(l1.Val)
	}
	for ; l2 != nil; l2 = l2.Next {
		num2.PushBack(l2.Val)
	}
	//max := math.Max(float64(len(num1)), float64(len(num2)))
	back1 := num1.Back()
	back2 := num2.Back()
	addOne := 0
	s := ""
	for back1 != nil || back2 != nil {
		var value1 = 0
		var value2 = 0
		if back1 != nil {
			value1 = back1.Value.(int)
		}
		if back2 != nil {
			value2 = back2.Value.(int)
		}
		i1 := value1
		i2 := value2
		if i1+i2+addOne >= 10 {
			s = strconv.Itoa((i1+i2+addOne)%10) + s
			addOne = 1
		} else {
			s = strconv.Itoa(i1+i2+addOne) + s
			addOne = 0
		}
		if back1 != nil {
			num1.Remove(back1)
		}
		if back2 != nil {
			num2.Remove(back2)
		}
		back1 = num1.Back()
		back2 = num2.Back()
	}
	if addOne == 1 {
		s = strconv.Itoa(addOne) + s
	}
	root := &ListNode{}
	tmp := root
	runes := []rune(s)
	for i := range runes {
		s := string(runes[i])
		atoi, _ := strconv.Atoi(s)
		tmp.Next = &ListNode{atoi, nil}
		tmp = tmp.Next
	}
	return root.Next
}
