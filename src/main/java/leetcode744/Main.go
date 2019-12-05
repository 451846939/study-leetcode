package main

import (
	"fmt"
)

/**
给定一个只包含小写字母的有序数组letters 和一个目标字母 target，寻找有序数组里面比目标字母大的最小字母。

数组里字母的顺序是循环的。举个例子，如果目标字母target = 'z' 并且有序数组为 letters = ['a', 'b']，则答案返回 'a'。

示例:

输入:
letters = ["c", "f", "j"]
target = "a"
输出: "c"

输入:
letters = ["c", "f", "j"]
target = "c"
输出: "f"

输入:
letters = ["c", "f", "j"]
target = "d"
输出: "f"

输入:
letters = ["c", "f", "j"]
target = "g"
输出: "j"

输入:
letters = ["c", "f", "j"]
target = "j"
输出: "c"

输入:
letters = ["c", "f", "j"]
target = "k"
输出: "c"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

	a := nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')
	fmt.Print(a)
}
func nextGreatestLetter(letters []byte, target byte) byte {
	var left = 0
	var right = len(letters)

	for left < right {
		var now = left + (right-left)>>1

		if target >= letters[now] {
			left = now + 1
		} else {
			right = now
		}
	}
	return letters[left%len(letters)]
}
