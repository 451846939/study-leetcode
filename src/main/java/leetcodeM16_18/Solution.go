package main

/*
你有两个字符串，即pattern和value。 pattern字符串由字母"a"和"b"组成，用于描述字符串中的模式。例如，字符串"catcatgocatgo"匹配模式"aabab"（其中"cat"是"a"，"go"是"b"），该字符串也匹配像"a"、"ab"和"b"这样的模式。但需注意"a"和"b"不能同时表示相同的字符串。编写一个方法判断value字符串是否匹配pattern字符串。

示例 1：

输入： pattern = "abba", value = "dogcatcatdog"
输出： true
示例 2：

输入： pattern = "abba", value = "dogcatcatfish"
输出： false
示例 3：

输入： pattern = "aaaa", value = "dogcatcatdog"
输出： false
示例 4：

输入： pattern = "abba", value = "dogdogdogdog"
输出： true
解释： "a"="dogdog",b=""，反之也符合规则
提示：

0 <= len(pattern) <= 1000
0 <= len(value) <= 1000
你可以假设pattern只包含字母"a"和"b"，value仅包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pattern-matching-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func patternMatching(pattern string, value string) bool {
	countA, countB := 0, 0
	for i := range pattern {
		if pattern[i] == 'a' {
			countA++
		} else {
			countB++
		}
	}
	if countA < countB {
		countA, countB = countB, countA
		tmp := ""
		for i := range pattern {
			if pattern[i] == 'a' {
				tmp += "b"
			} else {
				tmp += "a"
			}
		}
		pattern = tmp
	}
	if len(value) == 0 {
		return countB == 0
	}
	if len(pattern) == 0 {
		return false
	}
	for lenA := 0; countA*lenA <= len(value); lenA++ {
		rest := len(value) - countA*lenA
		if (countB == 0 && rest == 0) || (countB != 0 && rest%countB == 0) {
			lenB := -1
			if countB == 0 {
				lenB = 0
			} else {
				lenB = rest / countB
			}
			pos, correct := 0, true
			var valueA, valueB string
			for i := range pattern {
				if pattern[i] == 'a' {
					sub := value[pos : pos+lenA]
					if len(valueA) == 0 {
						valueA = sub
					} else if valueA != sub {
						correct = false
						break
					}
					pos += lenA
				} else {
					sub := value[pos : pos+lenB]
					if len(valueB) == 0 {
						valueB = sub
					} else if valueB != sub {
						correct = false
						break
					}
					pos += lenB
				}
			}
			if correct && valueA != valueB {
				return true
			}
		}
	}
	return false
}
