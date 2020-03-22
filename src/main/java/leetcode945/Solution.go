package main

import "fmt"

/*

给定整数数组 A，每次 move 操作将会选择任意 A[i]，并将其递增 1。

返回使 A 中的每个值都是唯一的最少操作次数。

示例 1:

输入：[1,2,2]
输出：1
解释：经过一次 move 操作，数组将变为 [1, 2, 3]。
示例 2:

输入：[3,2,1,2,1,7]
输出：6
解释：经过 6 次 move 操作，数组将变为 [3, 4, 1, 2, 5, 7]。
可以看出 5 次或 5 次以下的 move 操作是不能让数组的每个值唯一的。
提示：

0 <= A.length <= 40000
0 <= A[i] < 40000
https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/
*/
func main() {
	//fmt.Println(minIncrementForUnique([]int{3,2,1,2,1,7}))
	//fmt.Println(minIncrementForUnique([]int{0,0}))
	fmt.Println(minIncrementForUnique([]int{2, 2, 2, 1}))
	//fmt.Println(minIncrementForUnique([]int{1,1,2,0}))
}
func minIncrementForUnique(A []int) int {
	//hash路径压缩
	hash := make([]int, len(A)*2)
	for i := range hash {
		hash[i] = -1
	}
	count := 0
	for _, e := range A {
		index := 0
		if hash[e] == -1 {
			hash[e] = e
			index = e
		} else {
			i := e
			if hash[e] != e && hash[e] != -1 {
				i = hash[e]
			}
			for ; i < len(hash); i++ {
				if hash[i] == -1 {
					hash[i] = e
					index = i
					//路径压缩的根本
					hash[e] = i
					break
				}
			}
		}
		count += index - e
	}
	return count
}
