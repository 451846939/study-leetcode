package main

import "strconv"

/*

给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

*/
func main() {

}

func addStrings(num1 string, num2 string) string {
	needAdd := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || needAdd != 0; i, j = i-1, j-1 {
		l, r := 0, 0
		if i >= 0 {
			l = int(num1[i] - '0')
		}
		if j >= 0 {
			r = int(num2[j] - '0')
		}
		tmp := l + r + needAdd
		ans = strconv.Itoa(tmp%10) + ans
		needAdd = tmp / 10
	}
	return ans
}
