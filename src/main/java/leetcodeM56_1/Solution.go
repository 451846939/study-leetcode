package main

/*

一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。



示例 1：

输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
示例 2：

输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]


限制：

2 <= nums <= 10000
https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
*/
func main() {

}
func singleNumbers(nums []int) []int {
	//用于将所有的数异或起来
	k := 0

	for i := range nums {
		k ^= nums[i]
	}

	//获得k中最低位的1
	mask := 1

	//mask = k & (-k) 这种方法也可以得到mask
	for (k & mask) == 0 {
		mask <<= 1
	}

	a := 0
	b := 0
	for i := range nums {
		if (nums[i] & mask) == 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}
