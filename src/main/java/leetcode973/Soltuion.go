package main

import "math/rand"

/*
我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。

（这里，平面上两点之间的距离是欧几里德距离。）

你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。



示例 1：

输入：points = [[1,3],[-2,2]], K = 1
输出：[[-2,2]]
解释：
(1, 3) 和原点之间的距离为 sqrt(10)，
(-2, 2) 和原点之间的距离为 sqrt(8)，
由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。
示例 2：

输入：points = [[3,3],[5,-1],[-2,4]], K = 2
输出：[[3,3],[-2,4]]
（答案 [[-2,4],[3,3]] 也会被接受。）


提示：

1 <= K <= points.length <= 10000
-10000 < points[i][0] < 10000
-10000 < points[i][1] < 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/k-closest-points-to-origin
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	kClosest([][]int{{0, 1}, {1, 0}}, 2)
}
func less(p, q []int) bool {
	return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
}
func kClosest(points [][]int, k int) (ans [][]int) {
	rand.Shuffle(len(points), func(i, j int) {
		points[i], points[j] = points[j], points[i]
	})

	var quickSelect func(l, r int)
	quickSelect = func(l, r int) {
		if l == r {
			return
		}
		pivot := points[r]
		lessCount := l
		for i := l; i < r; i++ {
			if less(points[i], pivot) {
				points[i], points[lessCount] = points[lessCount], points[i]
				lessCount++
			}
		}
		points[r], points[lessCount] = points[lessCount], points[r]
		if lessCount+1 == k {
			return
		} else if lessCount+1 < k {
			quickSelect(lessCount+1, r)
		} else {
			quickSelect(l, lessCount-1)
		}
	}
	quickSelect(0, len(points)-1)
	return points[:k]
}
