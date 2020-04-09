package main

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func main() {

}
func generateParenthesis(n int) []string {
	res := make([]string, 0, n<<1)
	if n == 0 {
		return res
	}
	dfs("", 0, 0, n, &res)
	return res
}

func dfs(nows string, leftCount, rightCount, n int, res *[]string) {
	if leftCount == n && rightCount == n {
		*res = append(*res, nows)
		return
	}
	if leftCount < rightCount {
		return
	}
	if leftCount < n {
		dfs(nows+"(", leftCount+1, rightCount, n, res)
	}
	if rightCount < n {
		dfs(nows+")", leftCount, rightCount+1, n, res)
	}
}
