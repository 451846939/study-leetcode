package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

/*

在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(findKthLargest3([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest3([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, index int) int {
	q := randomPartition(nums, l, r)
	if q == index {
		return nums[q]
	} else if q < index {
		return quickSelect(nums, q+1, r, index)
	}
	return quickSelect(nums, l, q-1, index)
}

func randomPartition(nums []int, l int, r int) int {
	i := rand.Int()%(r-l+1) + l
	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, l, r)
}

func partition(nums []int, l int, r int) int {
	right := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= right {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

//--------------------------------------------------------

func findKthLargest2(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, size int) {
	for i := size / 2; i >= 0; i-- {
		maxHeapify(nums, i, size)
	}
}

func maxHeapify(nums []int, i int, size int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < size && nums[l] > nums[largest] {
		largest = l
	}
	if r < size && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, size)
	}
}

//--------------------------------------------------------
func findKthLargest3(nums []int, k int) int {
	h := make(IntHeap, 0, k)
	heap.Init(&h)
	for i := range nums {
		heap.Push(&h, nums[i])
		if h.Len() > k {
			heap.Pop(&h)
		}
	}
	return heap.Pop(&h).(int)
}

type IntHeap []int // 定义一个类型

func (h IntHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h IntHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h IntHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	m := x.(int)
	*h = append(*h, m)
}
