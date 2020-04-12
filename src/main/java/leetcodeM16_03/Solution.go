package main

import (
	"fmt"
	"math"
)

/*
给定两条线段（表示为起点start = {X1, Y1}和终点end = {X2, Y2}），如果它们有交点，请计算其交点，没有交点则返回空值。

要求浮点型误差不超过10^-6。若有多个交点（线段重叠）则返回 X 值最小的点，X 坐标相同则返回 Y 值最小的点。



示例 1：

输入：
line1 = {0, 0}, {1, 0}
line2 = {1, 1}, {0, -1}
输出： {0.5, 0}
示例 2：

输入：
line1 = {0, 0}, {3, 3}
line2 = {1, 1}, {2, 2}
输出： {1, 1}
示例 3：

输入：
line1 = {0, 0}, {1, 1}
line2 = {1, 0}, {2, 1}
输出： {}，两条线段没有交点


提示：

坐标绝对值不会超过 2^7
输入的坐标均是有效的二维坐标

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	//fmt.Println(intersection([]int{0,0},[]int{0,1},[]int{0,0},[]int{0,-1}))
	//fmt.Println(intersection([]int{0,0},[]int{1,0},[]int{1,1},[]int{0,-1}))

	fmt.Println(intersection([]int{-10, 48}, []int{-43, 46}, []int{-16, 59}, []int{-1, 85}))
}
func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	//矩阵解法
	x1 := start1[0]
	y1 := start1[1]
	x2 := end1[0]
	y2 := end1[1]
	x3 := start2[0]
	y3 := start2[1]
	x4 := end2[0]
	y4 := end2[1]
	//if (y4-y3)*(x2-x1) == (y2-y1)*(x4-x3) {
	//
	//}
	//(y2−y1)(x2−x)=(x2−x1)(y2−y)->Ax+By+C=0
	//A:=y1-y2 ,B:=x2-x1,C:=-(A*x2+B*y2)
	//Ax+By+C=0
	//Dx+Ey+F=0
	//		|-C B| 		|A B|
	//x0=	|	 | / 	|	|
	//		|-F	E|		|D E|

	//		|A -C| 		|A B|
	//y0=	|	 | / 	|	|
	//		|D -F|		|D E|
	A := y1 - y2
	B := x2 - x1
	C := -(A*x2 + B*y2)
	D := y3 - y4
	E := x4 - x3
	F := -(D*x4 + E*y4)

	if A*E-B*D != 0 {
		x0 := float64(B*F-C*E) / float64(A*E-B*D)
		y0 := float64(C*D-A*F) / float64(A*E-B*D)
		if math.Min(float64(x1), float64(x2)) > x0 || math.Max(float64(x1), float64(x2)) < x0 {
			return []float64{}
		}
		if math.Min(float64(x3), float64(x4)) > x0 || math.Max(float64(x3), float64(x4)) < x0 {
			return []float64{}
		}
		return []float64{x0, y0}
	}

	if C == 0 && F != 0 {
		return []float64{}
	}
	if C != 0 && F == 0 {
		return []float64{}
	}

	//if C != 0 && F != 0 && A*F - C*D != 0 {
	//	return []float64{}
	//}
	if math.Min(float64(x1), float64(x2)) > math.Max(float64(x3), float64(x4)) || math.Max(float64(x1), float64(x2)) < math.Min(float64(x3), float64(x4)) {
		return []float64{}
	}
	if math.Min(float64(y1), float64(y2)) > math.Max(float64(y3), float64(y4)) || math.Max(float64(y1), float64(y2)) < math.Min(float64(y3), float64(y4)) {
		return []float64{}
	}
	p1 := getres(x1, y1, x2, y2)
	p2 := getres(x3, y3, x4, y4)
	if p1[0] < p2[0] {
		return p2
	}
	return p1
}
func getres(x1, y1, x2, y2 int) []float64 {
	if x1 < x2 {
		return []float64{float64(x1), float64(y1)}
	} else if x1 > x2 {
		return []float64{float64(x2), float64(y2)}
	} else if y1 < y2 {
		return []float64{float64(x1), float64(y1)}
	}
	return []float64{float64(x2), float64(y2)}
}
func get(j [][]int) int {
	if len(j) != 2 {
		return 0
	}
	if len(j[0]) != 2 {
		return 0
	}
	return j[0][0]*j[1][1] - j[0][1]*j[1][0]
}
