package main

import "fmt"

/*
0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。

例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。



示例 1：

输入: n = 5, m = 3
输出: 3
示例 2：

输入: n = 10, m = 17
输出: 2


限制：

1 <= n <= 10^5
1 <= m <= 10^6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(lastRemaining2(5, 2))
}
func lastRemaining(n int, m int) int {
	/*
		第一轮是 [0, 1, 2, 3, 4] ，所以是 [0, 1, 2, 3, 4] 这个数组的多个复制。这一轮 2 删除了。

		第二轮开始时，从 3 开始，所以是 [3, 4, 0, 1] 这个数组的多个复制。这一轮 0 删除了。

		第三轮开始时，从 1 开始，所以是 [1, 3, 4] 这个数组的多个复制。这一轮 4 删除了。

		第四轮开始时，还是从 1 开始，所以是 [1, 3] 这个数组的多个复制。这一轮 1 删除了。

		最后剩下的数字是 3。

		图中的绿色的线指的是新的一轮的开头是怎么指定的，每次都是固定地向前移位 mm 个位置。

		然后我们从最后剩下的 3 倒着看，我们可以反向推出这个数字在之前每个轮次的位置。

		最后剩下的 3 的下标是 0。

		第四轮反推，补上 mm 个位置，然后模上当时的数组大小 22，位置是(0 + 3) % 2 = 1。

		第三轮反推，补上 mm 个位置，然后模上当时的数组大小 33，位置是(1 + 3) % 3 = 1。

		第二轮反推，补上 mm 个位置，然后模上当时的数组大小 44，位置是(1 + 3) % 4 = 0。

		第一轮反推，补上 mm 个位置，然后模上当时的数组大小 55，位置是(0 + 3) % 5 = 3。

		所以最终剩下的数字的下标就是3。因为数组是从0开始的，所以最终的答案就是3。

		总结一下反推的过程，就是 (当前index + m) % 上一轮剩余数字的个数。

		作者：sweetieeyi
		链接：https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/solution/javajie-jue-yue-se-fu-huan-wen-ti-gao-su-ni-wei-sh/
		来源：力扣（LeetCode）
		著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	*/
	f := 0
	for i := 2; i < n+1; i++ {
		f = (m + f) % i
	}
	return f
}

//超时
func lastRemaining2(n int, m int) int {
	ints := make([]int, n)
	for i := range ints {
		ints[i] = i
	}
	index := 0
	for len(ints) != 1 {
		index = (m - 1 + index) % len(ints)
		if index == 0 {
			ints = append(ints[index+1:])
		} else if index+1 > len(ints)-1 {
			ints = append(ints[:index])
		} else {
			ints = append(ints[:index], ints[index+1:]...)
		}
	}
	return ints[0]
}
