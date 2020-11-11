package main

import (
	"fmt"
	"math"
)

/*
视频游戏“辐射4”中，任务“通向自由”要求玩家到达名为“Freedom Trail Ring”的金属表盘，并使用表盘拼写特定关键词才能开门。

给定一个字符串 ring，表示刻在外环上的编码；给定另一个字符串 key，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。

最初，ring 的第一个字符与12:00方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00 方向对齐，然后按下中心按钮，以此逐个拼写完 key 中的所有字符。

旋转 ring 拼出 key 字符 key[i] 的阶段中：

您可以将 ring 顺时针或逆时针旋转一个位置，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00 方向对齐，并且这个字符必须等于字符 key[i] 。
如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key 的下一个字符（下一阶段）, 直至完成所有拼写。
示例：





输入: ring = "godding", key = "gd"
输出: 4
解释:
 对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
 对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
 当然, 我们还需要1步进行拼写。
 因此最终的输出是 4。
提示：

ring 和 key 的字符串长度取值范围均为 1 至 100；
两个字符串中都只有小写字符，并且均可能存在重复字符；
字符串 key 一定可以由字符串 ring 旋转拼出。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/freedom-trail
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(findRotateSteps("godding", "gd"))
	fmt.Println(findRotateSteps("iotfo", "fioot"))
}
func findRotateSteps(ring string, key string) int {
	cache := [26][]int{}
	for i, c := range ring {
		cache[c-'a'] = append(cache[c-'a'], i)
	}
	//定义 dp[i][j] 表示从前往后拼写出 key 的第 i 个字符，ring 的第 j 个字符与 12:00 方向对齐的最少步数（下标均从 0 开始）。

	dp := make([][]int, len(key))
	for i := range dp {
		dp[i] = make([]int, len(ring))
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	for _, p := range cache[key[0]-'a'] {
		dp[0][p] = min(p, len(ring)-p) + 1
	}
	for i := 1; i < len(key); i++ {
		for _, j := range cache[key[i]-'a'] {
			for _, k := range cache[key[i-1]-'a'] {
				//min(abs(j-k), len(ring)-abs(j-k))+1 表示在当前第 k 个字符与 12:00 方向对齐时第 j 个字符旋转到 12:00 方向并按下拼写的最少步数。
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), len(ring)-abs(j-k))+1)
			}
		}
	}
	return min(dp[len(key)-1]...)
}
func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
