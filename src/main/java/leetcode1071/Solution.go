package main

/*
对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。

返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。



示例 1：

输入：str1 = "ABCABC", str2 = "ABC"
输出："ABC"
示例 2：

输入：str1 = "ABABAB", str2 = "ABAB"
输出："AB"
示例 3：

输入：str1 = "LEET", str2 = "CODE"
输出：""


提示：

1 <= str1.length <= 1000
1 <= str2.length <= 1000
str1[i] 和 str2[i] 为大写英文字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/greatest-common-divisor-of-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

func gcdOfStrings(str1 string, str2 string) string {
	//其实看起来两个字符串之间能有这种神奇的关系是挺不容易的，我们希望能够找到一个简单的办法识别是否有解。
	//如果它们有公因子 abc，那么 str1 就是 mm 个 abc 的重复，str2 是 nn 个 abc 的重复，连起来就是 m+nm+n 个 abc，好像 m+nm+n 个 abc 跟 n+mn+m 个 abc 是一样的。
	//所以如果 str1 + str2 === str2 + str1 就意味着有解。
	//我们也很容易想到 str1 + str2 !== str2 + str1 也是无解的充要条件。
	//当确定有解的情况下，最优解是长度为 gcd(str1.length, str2.length) 的字符串。
	//这个理论最优长度是不是每次都能达到呢？是的。
	//因为如果能循环以它的约数为长度的字符串，自然也能够循环以它为长度的字符串，所以这个理论长度就是我们要找的最优解。
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[0:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
