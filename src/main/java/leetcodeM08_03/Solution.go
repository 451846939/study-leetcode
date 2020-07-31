package main

/*
魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。给定一个有序整数数组，编写一种方法找出魔术索引，若有的话，在数组A中找出一个魔术索引，如果没有，则返回-1。若有多个魔术索引，返回索引值最小的一个。

示例1:

 输入：nums = [0, 2, 3, 4, 5]
 输出：0
 说明: 0下标的元素为0
示例2:

 输入：nums = [1, 1, 1]
 输出：1
提示:

nums长度在[1, 1000000]之间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/magic-index-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func findMagicIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			return i
		}
	}
	return -1
}

func findMagicIndex2(nums []int) int {
	return getAns(nums, 0, len(nums)-1)
}

func getAns(nums []int, l int, r int) int {
	if l > r {
		return -1
	}
	mid := l + ((r - l) >> 1)
	lAns := getAns(nums, l, mid-1)
	if lAns != -1 {
		return lAns
	} else if nums[mid] == mid {
		return mid
	}
	return getAns(nums, mid+1, r)
}
