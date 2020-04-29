package main

/*
（这是一个 交互式问题 ）

给你一个 山脉数组 mountainArr，请你返回能够使得 mountainArr.get(index) 等于 target 最小 的下标 index 值。

如果不存在这样的下标 index，就请返回 -1。



何为山脉数组？如果数组 A 是一个山脉数组的话，那它满足如下条件：

首先，A.length >= 3

其次，在 0 < i < A.length - 1 条件下，存在 i 使得：

A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[A.length - 1]


你将 不能直接访问该山脉数组，必须通过 MountainArray 接口来获取数据：

MountainArray.get(k) - 会返回数组中索引为k 的元素（下标从 0 开始）
MountainArray.length() - 会返回该数组的长度


注意：

对 MountainArray.get 发起超过 100 次调用的提交将被视为错误答案。此外，任何试图规避判题系统的解决方案都将会导致比赛资格被取消。

为了帮助大家更好地理解交互式问题，我们准备了一个样例 “答案”：https://leetcode-cn.com/playground/RKhe3ave，请注意这 不是一个正确答案。



示例 1：

输入：array = [1,2,3,4,5,3,1], target = 3
输出：2
解释：3 在数组中出现了两次，下标分别为 2 和 5，我们返回最小的下标 2。
示例 2：

输入：array = [0,1,2,4,2,1], target = 3
输出：-1
解释：3 在数组中没有出现，返回 -1。


提示：

3 <= mountain_arr.length() <= 10000
0 <= target <= 10^9
0 <= mountain_arr.get(index) <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-in-mountain-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 */
type MountainArray struct {
}

func (this *MountainArray) get(index int) int {
	return 1
}
func (this *MountainArray) length() int {
	return 1
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	//先找山顶
	left, right := 0, mountainArr.length()-1
	for left != right {
		mid := (left + right) / 2
		if mountainArr.get(mid) > mountainArr.get(mid+1) {
			//右边有序，找左边
			right = mid
		} else {
			//左边有序，找右边
			left = mid + 1
		}
	}
	peak := left //结束后山顶就是left
	//先在左边找
	left, right = 0, peak
	for left != right {
		mid := (left + right) / 2
		val := mountainArr.get(mid)
		if val == target {
			return mid
		}
		if target > val {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if target == mountainArr.get(left) {
		return left
	}
	//再在右边找
	left, right = peak, mountainArr.length()-1
	for left != right {
		mid := (left + right) / 2
		val := mountainArr.get(mid)
		if val == target {
			return mid
		}
		if target < val {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if target == mountainArr.get(left) {
		return left
	}
	return -1
}
