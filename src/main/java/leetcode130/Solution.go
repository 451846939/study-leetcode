package main

/*
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X
解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

var (
	dx = [4]int{1, -1, 0, 0}
	dy = [4]int{0, 0, 1, -1}
)

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m := len(board), len(board[0])
	queue := [][]int{}
	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
		}
		if board[i][m-1] == 'O' {
			queue = append(queue, []int{i, m - 1})
		}
	}
	for i := 1; i < m-1; i++ {
		if board[0][i] == 'O' {
			queue = append(queue, []int{0, i})
		}
		if board[n-1][i] == 'O' {
			queue = append(queue, []int{n - 1, i})
		}
	}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]
		board[x][y] = 'A'
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx < 0 || my < 0 || mx >= n || my >= m || board[mx][my] != 'O' {
				continue
			}
			queue = append(queue, []int{mx, my})
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
