package main

import "fmt"

/*
给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。



示例：

输入：A = [4,5,0,-2,-3,1], K = 5
输出：7
解释：
有 7 个子数组满足其元素之和可被 K = 5 整除：
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]


提示：

1 <= A.length <= 30000
-10000 <= A[i] <= 10000
2 <= K <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarray-sums-divisible-by-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(-8 % 5)
}
func subarraysDivByK(A []int, K int) int {
	sum := 0
	mod := make([]int, K)
	ans := 0
	mod[0] = 1
	for i := range A {
		sum += A[i]
		m := sum % K
		if m < 0 {
			m += K
		}
		//边存边查看 map ，如果 map 中已存在 key 为 当前前缀和 mod K
		//说明存在 之前求过的某个前缀和，它 mod K == 当前前缀和 mod K
		//把满足条件的 key 对应的出现次数，累加给 ans
		ans += mod[m]
		mod[m]++
	}
	return ans
}
