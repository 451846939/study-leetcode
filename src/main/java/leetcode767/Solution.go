package main

import "fmt"

/*
给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。

若可行，输出任意可行的结果。若不可行，返回空字符串。

示例 1:

输入: S = "aab"
输出: "aba"
示例 2:

输入: S = "aaab"
输出: ""
注意:

S 只包含小写字母并且长度在[1, 500]区间内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorganize-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(reorganizeString("aab"))
}
func reorganizeString(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}

	cnt := [26]int{}
	maxCnt := 0
	for _, ch := range S {
		ch -= 'a'
		cnt[ch]++
		if cnt[ch] > maxCnt {
			maxCnt = cnt[ch]
		}
	}
	if maxCnt > (n+1)/2 {
		return ""
	}
	ans := make([]byte, n)
	evenIndex, oddIndex, halfLen := 0, 1, n>>1
	for i := range cnt {
		ch := byte('a' + i)
		for cnt[i] > 0 && cnt[i] <= halfLen && oddIndex < n {
			ans[oddIndex] = ch
			cnt[i]--
			oddIndex += 2
		}
		for cnt[i] > 0 {
			ans[evenIndex] = ch
			cnt[i]--
			evenIndex += 2
		}
	}
	return string(ans)
}
