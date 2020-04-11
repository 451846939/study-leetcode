package main

import "fmt"

/*
你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。

每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。

你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。

每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。

你的目标是确切地知道 F 的值是多少。

无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？



示例 1：

输入：K = 1, N = 2
输出：2
解释：
鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
如果它没碎，那么我们肯定知道 F = 2 。
因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
示例 2：

输入：K = 2, N = 6
输出：3
示例 3：

输入：K = 3, N = 14
输出：4


提示：

1 <= K <= 100
1 <= N <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/super-egg-drop
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(superEggDrop3(2, 6))
}
func superEggDrop(K int, N int) int {
	if K < 1 || N < 1 {
		return 0
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
	}
	i := 0
	for dp[i][K] < N {
		i++
		for j := 1; j <= K; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j] + 1
		}
	}
	return i
}
func superEggDrop3(K int, N int) int {
	if K < 1 || N < 1 {
		return 0
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
		for j := range dp[i] {
			dp[i][j] = i
		}
	}

	dp[1][0] = 0
	for j := 1; j <= K; j++ {
		dp[1][j] = 1
	}
	for i := 0; i <= N; i++ {
		dp[i][0] = 0
		dp[i][1] = i
	}

	for i := 2; i <= N; i++ {
		for j := 2; j <= K; j++ {
			l := 1
			r := i
			for l < r {
				mid := l + ((r - l) + 1>>1)
				bork := dp[mid-1][j-1]
				noBork := dp[i-mid][j]
				if bork > noBork {
					r = mid - 1
				} else {
					l = mid
				}
			}
			dp[i][j] = max(dp[l-1][j-1], dp[i-l][j]) + 1
		}
	}
	return dp[N][K]
}
func superEggDrop2(K int, N int) int {
	if K < 1 || N < 1 {
		return 0
	}
	dp := make([][]int, K+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
		for j := range dp[i] {
			if i == 0 || j == 0 {
				continue
			}
			dp[i][j] = j
		}
	}
	for i := 2; i <= K; i++ {
		for j := 1; j <= N; j++ {
			for k := 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], max(dp[i-1][k-1], dp[i][j-k])+1)
			}
		}
	}
	return dp[K][N]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
