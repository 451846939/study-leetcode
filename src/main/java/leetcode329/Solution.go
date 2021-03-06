package main

/*
给定一个整数矩阵，找出最长递增路径的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。

示例 1:

输入: nums =
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
输出: 4
解释: 最长递增路径为 [1, 2, 6, 9]。
示例 2:

输入: nums =
[
  [3,4,5],
  [3,2,6],
  [2,2,1]
]
输出: 4
解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

var (
	dirs          = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	rows, columns int
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows = len(matrix)
	columns = len(matrix[0])
	outdegrees := make([][]int, rows)
	for i := 0; i < rows; i++ {
		outdegrees[i] = make([]int, columns)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			for _, dir := range dirs {
				newRow, newColumn := i+dir[0], j+dir[1]
				if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[i][j] {
					outdegrees[i][j]++
				}
			}
		}
	}
	q := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if outdegrees[i][j] == 0 {
				q = append(q, []int{i, j})
			}
		}
	}
	ans := 0
	for len(q) != 0 {
		ans++
		size := len(q)
		for i := 0; i < size; i++ {
			cell := q[0]
			q = q[1:]
			r, c := cell[0], cell[1]
			for _, dir := range dirs {
				newRow, newColumn := r+dir[0], c+dir[1]
				if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] < matrix[r][c] {
					outdegrees[newRow][newColumn]--
					if outdegrees[newRow][newColumn] == 0 {
						q = append(q, []int{newRow, newColumn})
					}
				}
			}
		}
	}
	return ans
}
