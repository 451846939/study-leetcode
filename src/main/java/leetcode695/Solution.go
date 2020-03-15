package main

import "fmt"

/*给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

示例 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

示例 2:

[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。

注意: 给定的矩阵grid 的长度和宽度都不超过 50。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-area-of-island
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(maxAreaOfIsland([][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}))
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
	fmt.Println(maxAreaOfIsland([][]int{{1}}))
}

type point struct {
	x, y int
}

var q = make([]point, 0)

func next(Point point, points [][]int) (bool, int) {
	count := 0
	flag := false
	//        left
	if Point.x-1 >= 0 && Point.x-1 < len(points) && points[Point.x-1][Point.y] == 1 {
		newPoint := point{0, 0}
		newPoint.x = Point.x - 1
		newPoint.y = Point.y
		points[newPoint.x][newPoint.y] = 2
		points[Point.x][Point.y] = 2
		q = append(q, newPoint)
		flag = true
		count++
	}
	//        up
	if Point.y-1 >= 0 && Point.y-1 < len(points[0]) && points[Point.x][Point.y-1] == 1 {
		newPoint := point{0, 0}
		newPoint.x = Point.x
		newPoint.y = Point.y - 1
		points[newPoint.x][newPoint.y] = 2
		points[Point.x][Point.y] = 2
		q = append(q, newPoint)
		flag = true
		count++
	}
	//        right
	if Point.x+1 >= 0 && Point.x+1 < len(points) && points[Point.x+1][Point.y] == 1 {
		newPoint := point{0, 0}
		newPoint.x = Point.x + 1
		newPoint.y = Point.y
		points[newPoint.x][newPoint.y] = 2
		points[Point.x][Point.y] = 2
		q = append(q, newPoint)
		flag = true
		count++
	}
	//        down
	if Point.y+1 >= 0 &&
		Point.y+1 < len(points[0]) &&
		points[Point.x][Point.y+1] == 1 {
		newPoint := point{0, 0}
		newPoint.x = Point.x
		newPoint.y = Point.y + 1
		points[newPoint.x][newPoint.y] = 2
		points[Point.x][Point.y] = 2
		q = append(q, newPoint)
		flag = true
		count++
	}
	return flag, count
}

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				q = append(q, point{i, j})
				//找到一个点后开始广度遍历
				tmp := 1
				for len(q) != 0 {
					nowPoint := q[0]
					flag, count := next(nowPoint, grid)
					if flag {
						tmp += count
					}
					res = max(res, tmp)
					q = q[1:]
				}
			}
		}
	}
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
