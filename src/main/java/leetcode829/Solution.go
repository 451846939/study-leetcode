package main

import "fmt"

/**

给定一个正整数 N，试求有多少组连续正整数满足所有数字之和为 N?

示例 1:

输入: 5
输出: 2
解释: 5 = 5 = 2 + 3，共有两组连续整数([5],[2,3])求和后为 5。
示例 2:

输入: 9
输出: 3
解释: 9 = 9 = 4 + 5 = 2 + 3 + 4
示例 3:

输入: 15
输出: 4
解释: 15 = 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5
说明: 1 <= N <= 10 ^ 9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/consecutive-numbers-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func main() {
	fmt.Println(consecutiveNumbersSum(246854111))
}
func consecutiveNumbersSum(N int) int {
	res := 0
	left := 1
	right := 1
	sum := 0
	for right <= N {
		if sum == N {
			res++
			sum -= left
			left++
		} else if sum > N {
			sum -= left
			left++
		} else {
			sum += right
			right++
		}
	}
	return res + 1
}
func consecutiveNumbersSum2(N int) int {
	for (N & 1) == 0 {
		N >>= 1
	}
	ans := 1
	d := 3

	for d*d <= N {
		e := 0
		for N%d == 0 {
			N /= d
			e++
		}
		ans *= e + 1
		d += 2
	}

	if N > 1 {
		ans <<= 1
	}
	return ans
}
