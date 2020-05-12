package main

import "fmt"

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。


示例:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Pop()
	param3 := obj.Top()
	param4 := obj.GetMin()
	fmt.Println(param3, param4)
}

type MinStack struct {
	data   *[]int
	helper *[]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	data := make([]int, 0)
	helper := make([]int, 0)
	return MinStack{data: &data, helper: &helper}
}

func (this *MinStack) Push(x int) {
	data := this.data
	helper := this.helper
	*data = append(*data, x)
	if len(*helper) == 0 || (*helper)[len(*helper)-1] >= x {
		*helper = append(*helper, x)
	}
}

func (this *MinStack) Pop() {
	data := this.data
	helper := this.helper
	if (*data)[len(*data)-1] == (*helper)[len(*helper)-1] {
		*data = (*data)[:len(*data)-1]
		*helper = (*helper)[:len(*helper)-1]
	} else {
		*data = (*data)[:len(*data)-1]
	}
}

func (this *MinStack) Top() int {
	data := this.data
	if len(*data) != 0 {
		return (*data)[len(*data)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	helper := this.helper
	if len(*helper) != 0 {
		return (*helper)[len(*helper)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
