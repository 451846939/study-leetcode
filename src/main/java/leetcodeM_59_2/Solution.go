package main

import "fmt"

/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]


限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	obj := Constructor()
	param_1 := obj.Max_value()
	obj.Push_back(1)
	obj.Push_back(3)
	obj.Push_back(3)
	_ = obj.Pop_front()
	_ = obj.Pop_front()
	_ = obj.Pop_front()
	_ = obj.Pop_front()
	_ = obj.Max_value()
	_ = obj.Pop_front()
	obj.Push_back(2)
	param_3 := obj.Pop_front()

	fmt.Println(obj.Max_value())
	fmt.Println(param_1)
	fmt.Println(param_3)
}

type MaxQueue struct {
	size     int
	queue    []int
	maxQueue []int
}

func Constructor() MaxQueue {
	var a = MaxQueue{0, make([]int, 0), make([]int, 0)}
	return a
}

func (this *MaxQueue) Max_value() int {
	if this.size == 0 {
		return -1
	}
	return this.maxQueue[0]
}

func (this *MaxQueue) Push_back(value int) {
	for index := len(this.maxQueue) - 1; index >= 0; index-- {
		if len(this.maxQueue) > 0 && this.maxQueue[index] < value {
			this.maxQueue = this.maxQueue[:index]
		}
	}
	this.maxQueue = append(this.maxQueue, value)
	this.queue = append(this.queue, value)
	this.size += 1
}

func (this *MaxQueue) Pop_front() int {
	if this.size == 0 {
		return -1
	}
	head := this.queue[0]
	this.size -= 1
	if this.maxQueue[0] == head {
		this.maxQueue = this.maxQueue[1:]
	}
	this.queue = this.queue[1:]
	return head
}
