package main

import "fmt"

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
