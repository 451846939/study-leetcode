package main

/*
给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个子序列，其中每个子序列都由连续整数组成且长度至少为 3 。

如果可以完成上述分割，则返回 true ；否则，返回 false 。



示例 1：

输入: [1,2,3,3,4,5]
输出: True
解释:
你可以分割出这样两个连续子序列 :
1, 2, 3
3, 4, 5


示例 2：

输入: [1,2,3,3,4,4,5,5]
输出: True
解释:
你可以分割出这样两个连续子序列 :
1, 2, 3, 4, 5
3, 4, 5


示例 3：

输入: [1,2,3,4,4,5]
输出: False


提示：

输入的数组长度范围为 [1, 10000]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func isPossible(nums []int) bool {
	left := make(map[int]int)
	for _, v := range nums {
		left[v]++
	}
	endCount := make(map[int]int)
	for _, v := range nums {
		if left[v] == 0 {
			continue
		}
		if endCount[v-1] > 0 {
			left[v]--
			endCount[v-1]--
			endCount[v]++
		} else if left[v+1] > 0 && left[v+2] > 0 {
			left[v]--
			left[v+1]--
			left[v+2]--
			endCount[v+2]++
		} else {
			return false
		}
	}
	return true
}
