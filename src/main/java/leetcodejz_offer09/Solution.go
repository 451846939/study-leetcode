package main

import "container/list"

/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )



示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

//type CQueue struct {
//	q []int
//}
//
//
//func Constructor() CQueue {
//	return CQueue{make([]int,0)}
//}
//
//
//func (this *CQueue) AppendTail(value int)  {
//	this.q=append(this.q, value)
//}
//
//
//func (this *CQueue) DeleteHead() int {
//	if len(this.q)==0 {
//		return -1
//	}
//	i := this.q[0]
//	this.q=this.q[1:]
//	return i
//}

type CQueue struct {
	s1, s2 *list.List
}

func Constructor() CQueue {
	return CQueue{list.New(), list.New()}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.s2.Len() != 0 {
		remove := this.s2.Remove(this.s2.Back())
		return remove.(int)
	} else {
		for this.s1.Len() != 0 {
			this.s2.PushBack(this.s1.Remove(this.s1.Back()))
		}
	}
	if this.s2.Len() == 0 {
		return -1
	}
	return this.s2.Remove(this.s2.Back()).(int)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
