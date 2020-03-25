package main

/*

在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。

请你返回最终形体的表面积。



示例 1：

输入：[[2]]
输出：10
示例 2：

输入：[[1,2],[3,4]]
输出：34
示例 3：

输入：[[1,0],[0,2]]
输出：16
示例 4：

输入：[[1,1,1],[1,0,1],[1,1,1]]
输出：32
示例 5：

输入：[[2,2,2],[2,1,2],[2,2,2]]
输出：46


提示：

1 <= N <= 50
0 <= grid[i][j] <= 50
https://leetcode-cn.com/problems/surface-area-of-3d-shapes/
*/
func main() {

}
func surfaceArea(grid [][]int) int {
	count := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 0 {
				continue
			}
			//计算当前z轴的表面积
			count += grid[x][y]*6 - grid[x][y]*2 + 2
			//扫描之前的x轴扫描是否有重叠
			if x-1 >= 0 && grid[x-1][y] != 0 {
				count -= min(grid[x][y], grid[x-1][y]) * 2
			}
			//扫描之前的y轴扫描是否有重叠
			if y-1 >= 0 && grid[x][y-1] != 0 {
				count -= min(grid[x][y], grid[x][y-1]) * 2
			}
		}
	}
	return count
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
