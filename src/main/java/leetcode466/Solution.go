package main

/*
由 n 个连接的字符串 s 组成字符串 S，记作 S = [s,n]。例如，["abc",3]=“abcabcabc”。

如果我们可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。例如，根据定义，"abc" 可以从 “abdbec” 获得，但不能从 “acbbe” 获得。

现在给你两个非空字符串 s1 和 s2（每个最多 100 个字符长）和两个整数 0 ≤ n1 ≤ 106 和 1 ≤ n2 ≤ 106。现在考虑字符串 S1 和 S2，其中 S1=[s1,n1] 、S2=[s2,n2] 。

请你找出一个可以满足使[S2,M] 从 S1 获得的最大整数 M 。



示例：

输入：
s1 ="acb",n1 = 4
s2 ="ab",n2 = 2

返回：
2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-the-repetitions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {

	len1, len2 := len(s1), len(s2)
	index1, index2 := 0, 0 // 注意此处直接使用的下标，不取模

	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return 0
	}

	map1, map2 := make(map[int]int), make(map[int]int)
	ans := 0 // 注意，此处存储的是 n1个s1  中 s2 的个数，而非 n1个s1  中 n2个s2 的个数
	for index1/len1 < n1 {
		if index1%len1 == len1-1 {
			if val, ok := map1[index2%len2]; ok {
				//	出现循环，开始快进

				cycleLen := index1/len1 - val/len1
				// 还有多少个循环
				cycleNum := (n1 - 1 - index1/len1) / cycleLen

				cycleS2Num := index2/len2 - map2[index2%len2]/len2

				index1 += cycleNum * cycleLen * len1

				ans += cycleNum * cycleS2Num

			} else {
				map1[index2%len2] = index1
				map2[index2%len2] = index2
			}
		}
		if s1[index1%len1] == s2[index2%len2] {
			if index2%len2 == len2-1 {
				ans++
			}
			index2++
		}
		index1++
	}
	return ans / n2
}
