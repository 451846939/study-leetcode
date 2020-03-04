package main

/**
 * 在给定的网格中，每个单元格可以有以下三个值之一：
 *
 * 值 0 代表空单元格；
 * 值 1 代表新鲜橘子；
 * 值 2 代表腐烂的橘子。
 * 每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
 *
 * 返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。
 *https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/02/16/oranges.png
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：[[2,1,1],[1,1,0],[0,1,1]]
 * 输出：4
 * 示例 2：
 *
 * 输入：[[2,1,1],[0,1,1],[1,0,1]]
 * 输出：-1
 * 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
 * 示例 3：
 *
 * 输入：[[0,2]]
 * 输出：0
 * 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
 *
 *
 * 提示：
 *
 * 1 <= grid.length <= 10
 * 1 <= grid[0].length <= 10
 * grid[i][j] 仅为 0、1 或 2
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/rotting-oranges
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func main() {
	orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
}

type point struct {
	x, y, count int
}

var bedPoint = make([]point, 0)
var points [][]int

func next(Point point) bool {
	flag := false
	//        left
	if Point.x-1 >= 0 && Point.x-1 < len(points) && points[Point.x-1][Point.y] == 1 {
		//            count++;
		newPoint := point{0, 0, 0}
		newPoint.x = Point.x - 1
		newPoint.y = Point.y
		points[newPoint.x][newPoint.y] = 2
		newPoint.count = Point.count + 1
		bedPoint = append(bedPoint, newPoint)
		flag = true
	}
	//        up
	if Point.y-1 >= 0 && Point.y-1 < len(points[0]) && points[Point.x][Point.y-1] == 1 {
		//            count++;
		newPoint := point{0, 0, 0}
		newPoint.x = Point.x
		newPoint.y = Point.y - 1
		points[newPoint.x][newPoint.y] = 2
		newPoint.count = Point.count + 1
		bedPoint = append(bedPoint, newPoint)
		flag = true
	}
	//        right
	if Point.x+1 >= 0 && Point.x+1 < len(points) && points[Point.x+1][Point.y] == 1 {
		//            count++;
		newPoint := point{0, 0, 0}
		newPoint.x = Point.x + 1
		newPoint.y = Point.y
		points[newPoint.x][newPoint.y] = 2
		newPoint.count = Point.count + 1
		bedPoint = append(bedPoint, newPoint)
		flag = true
	}
	//        down
	if Point.y+1 >= 0 &&
		Point.y+1 < len(points[0]) &&
		points[Point.x][Point.y+1] == 1 {
		//            count++;
		newPoint := point{0, 0, 0}
		newPoint.x = Point.x
		newPoint.y = Point.y + 1
		points[newPoint.x][newPoint.y] = 2
		newPoint.count = Point.count + 1
		bedPoint = append(bedPoint, newPoint)
		flag = true
	}
	return flag
}
func orangesRotting(grid [][]int) int {
	points = grid
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				bedPoint = append(bedPoint, point{i, j, 0})
			}
		}
	}
	for len(bedPoint) > 0 {
		curr := bedPoint[0]
		bedPoint = bedPoint[1:]
		next(curr)
		count = curr.count
	}
	for i := range points {
		for j := range points[i] {
			if points[i][j] == 1 {
				return -1
			}
		}
	}
	return count
}
