package main

/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。



示例 1:

输入: [7,5,6,4]
输出: 5


限制：

0 <= 数组长度 <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + ((end - start) >> 1)
	count := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	tmp := make([]int, 0)
	l, r := start, mid+1
	for l <= mid && r <= end {
		if nums[l] <= nums[r] {
			tmp = append(tmp, nums[l])
			l++
		} else {
			tmp = append(tmp, nums[r])
			r++
			count += mid - l + 1
		}
	}
	for ; l <= mid; l++ {
		tmp = append(tmp, nums[l])
		//count += end - (mid + 1) + 1
	}
	for ; r <= end; r++ {
		tmp = append(tmp, nums[r])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return count
}
