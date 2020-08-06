package main

/*
给定一组唯一的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。

示例 1:

输入: ["abcd","dcba","lls","s","sssll"]
输出: [[0,1],[1,0],[3,2],[2,4]]
解释: 可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
示例 2:

输入: ["bat","tab","cat"]
输出: [[0,1],[1,0]]
解释: 可拼接成的回文串为 ["battab","tabbat"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

type Node struct {
	ch   [26]int
	flag int
}

var tree []Node

func palindromePairs(words []string) [][]int {
	tree = []Node{{[26]int{}, -1}}
	n := len(words)
	for i := 0; i < n; i++ {
		insert(words[i], i)
	}
	ret := [][]int{}
	for i := 0; i < n; i++ {
		word := words[i]
		m := len(word)
		for j := 0; j <= m; j++ {
			if isPalindrome(word, j, m-1) {
				leftId := findWord(word, 0, j-1)
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word, 0, j-1) {
				rightId := findWord(word, j, m-1)
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{rightId, i})
				}
			}
		}
	}
	return ret
}
func insert(s string, id int) {
	add := 0
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		if tree[add].ch[x] == 0 {
			tree = append(tree, Node{[26]int{}, -1})
			tree[add].ch[x] = len(tree) - 1
		}
		add = tree[add].ch[x]
	}
	tree[add].flag = id
}

func findWord(s string, left, right int) int {
	add := 0
	for i := right; i >= left; i-- {
		x := int(s[i] - 'a')
		if tree[add].ch[x] == 0 {
			return -1
		}
		add = tree[add].ch[x]
	}
	return tree[add].flag
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right-left+1)>>1; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}
