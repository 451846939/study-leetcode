package main

import "fmt"

/*
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。

示例 1:
输入:

0 0 0
0 1 0
0 0 0
输出:

0 0 0
0 1 0
0 0 0
示例 2:
输入:

0 0 0
0 1 0
1 1 1
输出:

0 0 0
0 1 0
1 2 1
注意:

给定矩阵的元素个数不超过 10000。
给定矩阵中至少有一个元素是 0。
矩阵中的元素只在四个方向上相邻: 上、下、左、右。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/01-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	//[[0,0,0],[0,1,0],[1,1,1]]
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}
func updateMatrix(matrix [][]int) [][]int {
	seen := make([][]int, len(matrix))
	q := make([]point, 0, len(matrix))
	for x := range matrix {
		seen[x] = make([]int, len(matrix[x]))
		for y := range matrix[x] {
			if matrix[x][y] == 0 {
				q = append(q, point{x, y})
				seen[x][y] = 1
			}
		}
	}
	dicx := []int{0, 1, 0, -1}
	dicy := []int{1, 0, -1, 0}
	for len(q) != 0 {
		now := q[0]
		for i := 0; i < 4; i++ {
			nx := now.x + dicx[i]
			ny := now.y + dicy[i]
			if nx >= len(matrix) || nx < 0 || ny >= len(matrix[0]) || ny < 0 || seen[nx][ny] != 0 {
				continue
			}
			matrix[nx][ny] = matrix[now.x][now.y] + 1
			q = append(q, point{nx, ny})
			seen[nx][ny] = 1
		}
		q = q[1:]
	}
	return matrix
}

type point struct {
	x, y int
}
