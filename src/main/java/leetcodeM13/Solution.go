package main

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？



示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 1：

输入：m = 3, n = 1, k = 0
输出：1
提示：

1 <= n,m <= 100
0 <= k <= 20

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

type point struct {
	x, y int
}

func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	maps := make([][]int, m)
	for i := range maps {
		maps[i] = make([]int, n)
		for j := range maps[i] {
			maps[i][j] = 0
		}
	}
	mx := []int{1, 0, -1, 0}
	my := []int{0, 1, 0, -1}
	q := make([]point, 0, m)
	q = append(q, point{0, 0})
	count := 1
	for len(q) != 0 {
		nowp := q[0]
		maps[nowp.x][nowp.y] = 1
		for i := 0; i < 2; i++ {
			newx := nowp.x + mx[i]
			newy := nowp.y + my[i]
			if newx >= 0 && newx < m && newy >= 0 && newy < n && maps[newx][newy] != 1 && (get(newx)+get(newy) <= k) {
				q = append(q, point{newx, newy})
				maps[newx][newy] = 1
				count++
			}
		}
		q = q[1:]
	}
	return count
}

// 计算 x 的数位之和
func get(x int) int {
	res := 0
	for ; x != 0; x /= 10 {
		res += x % 10
	}
	return res
}
