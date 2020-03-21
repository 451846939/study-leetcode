package main

import (
	"fmt"
)

/*

有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？

如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。

你允许：

装满任意一个水壶
清空任意一个水壶
从一个水壶向另外一个水壶倒水，直到装满或者倒空
示例 1: (From the famous "Die Hard" example)

输入: x = 3, y = 5, z = 4
输出: True
示例 2:

输入: x = 2, y = 6, z = 5
输出: False
https://leetcode-cn.com/problems/water-and-jug-problem/
*/
func main() {
	//fmt.Println(gcd(3,5))
	//fmt.Println(canMeasureWater2(3,5,4))
	fmt.Println(canMeasureWater3(0, 2, 1))
	fmt.Println(canMeasureWater3(1, 2, 3))
}

type point struct {
	x, y int
}

func canMeasureWater3(x int, y int, z int) bool {
	//图的遍历（x,y）表示状态，这里使用了简化的遍历，偷了个懒，实际应该简历坐标系遍历
	if z < 0 || z > x+y {
		return false
	}
	cache := make(map[int]int)
	q := make([]int, 1)
	q[0] = 0
	for len(q) != 0 {
		tmp := q[0]
		q = q[1:]
		if tmp <= y && cache[tmp+x] != 1 {
			cache[tmp+x] = 1
			q = append(q, tmp+x)
		}
		if tmp <= x && cache[tmp+y] != 1 {
			cache[tmp+y] = 1
			q = append(q, tmp+y)
		}
		if tmp-x >= 0 && cache[tmp-x] != 1 {
			cache[tmp-x] = 1
			q = append(q, tmp-x)
		}
		if tmp-y >= 0 && cache[tmp-y] != 1 {
			cache[tmp-y] = 1
			q = append(q, tmp-y)
		}
		if cache[z] == 1 {
			return true
		}
	}
	return false
}

func canMeasureWater2(x int, y int, z int) bool {
	//最大公约数做法
	if x == z || y == z || z == 0 {
		return true
	}
	if x == y || x+y < z {
		return false
	}
	i := gcd3(x, y)
	if x < y {
		return z%i == 0
	} else {
		return z%i == 0
	}
}

func gcd3(x, y int) int {
	f := func(x, y int) (int, int) {
		if x == 0 {
			return y, 0
		}
		if y == 0 {
			return x, 0
		}
		if y > x {
			return x, y - x
		}
		return y, x - y
	}
	res := 0
	newx, newy := f(x, y)
	for newx != newy && newx != 0 && newy != 0 {
		newx, newy = f(newx, newy)
	}
	if newy == 0 {
		res = newx
	} else {
		res = newy
	}
	return res
}
func gcd2(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	if y == x {
		return x
	}
	if y > x {
		return gcd2(x, y-x)
	}
	return gcd2(y, x-y)
}
func canMeasureWater(x int, y int, z int) bool {
	if x == z || y == z || z == 0 {
		return true
	}
	if x == y || x+y < z {
		return false
	}

	return z%gcd(x, y) == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
