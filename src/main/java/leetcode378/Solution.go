package main

/*
给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。



示例：

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

返回 13。


提示：
你可以假设 k 的值永远是有效的，1 ≤ k ≤ n2 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		mid := l + ((r - l) >> 1)
		if count(matrix, mid, n) >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func count(matrix [][]int, mid, n int) int {
	i, j := n-1, 0
	count := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			count += i + 1
			j++
		} else {
			i--
		}
	}
	return count
}
