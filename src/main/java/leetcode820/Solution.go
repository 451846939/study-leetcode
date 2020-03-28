package main

import (
	"fmt"
	"sort"
)

/*
给定一个单词列表，我们将这个列表编码成一个索引字符串 S 与一个索引列表 A。

例如，如果这个列表是 ["time", "me", "bell"]，我们就可以将其表示为 S = "time#bell#" 和 indexes = [0, 2, 5]。

对于每一个索引，我们可以通过从字符串 S 中索引的位置开始读取字符串，直到 "#" 结束，来恢复我们之前的单词列表。

那么成功对给定单词列表进行编码的最小字符串长度是多少呢？



示例：

输入: words = ["time", "me", "bell"]
输出: 10
说明: S = "time#bell#" ， indexes = [0, 2, 5] 。


提示：

1 <= words.length <= 2000
1 <= words[i].length <= 7
每个单词都是小写字母 。
https://leetcode-cn.com/problems/short-encoding-of-words/
*/
func main() {
	fmt.Println(minimumLengthEncoding2([]string{"time", "me", "bell"}))
}
func minimumLengthEncoding(words []string) int {
	hashmap := make(map[string]int, len(words))
	for _, e := range words {
		hashmap[e]++
	}

	for _, e := range words {
		for i := len(e) - 1; i >= 1; i-- {
			if hashmap[e[i:]] > 0 {
				hashmap[e[i:]] = 0
			}
		}
	}
	count := 0
	for k, v := range hashmap {
		if v != 0 {
			count += len(k) + 1
		}
	}

	return count
}

func minimumLengthEncoding2(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	constructor := Constructor()
	count := 0
	for i := range words {
		count += constructor.insert(words[i])
	}
	return count
}

type Trie struct {
	trieNode *TrieNode
}

func Constructor() Trie {
	return Trie{&TrieNode{'/', make([]*TrieNode, 26)}}
}

type TrieNode struct {
	data     uint8
	children []*TrieNode
}

func (this *Trie) insert(word string) int {
	p := this.trieNode
	isnew := false
	for i := len(word) - 1; i >= 0; i-- {
		index := word[i] - 'a'
		if p.children[index] == nil {
			isnew = true
			p.children[index] = &TrieNode{word[i], make([]*TrieNode, 26)}
		}
		p = p.children[index]
	}
	if isnew {
		return len(word) + 1
	}
	return 0
}
