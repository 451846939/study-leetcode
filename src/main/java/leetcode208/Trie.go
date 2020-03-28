package main

func main() {

}

type Trie struct {
	TrieNode *TrieNode
}
type TrieNode struct {
	data         uint8
	children     []*TrieNode
	isEndingChar bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{&TrieNode{'/', make([]*TrieNode, 26), false}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this.TrieNode
	for i := range word {
		index := word[i] - 'a'
		if p.children[index] == nil {
			p.children[index] = &TrieNode{word[i], make([]*TrieNode, 26), false}
		}
		p = p.children[index]
	}
	p.isEndingChar = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this.TrieNode
	for i := range word {
		index := word[i] - 'a'
		if p.children[index] == nil {
			return false
		}
		p = p.children[index]
	}
	if p.isEndingChar == false {
		return false
	} else {
		return true
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this.TrieNode
	for i := range prefix {
		index := prefix[i] - 'a'
		if p.children[index] == nil {
			return false
		}
		p = p.children[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
