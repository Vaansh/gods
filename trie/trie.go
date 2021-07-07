package trie

const ALPHABET = 26

type Node struct {
	children [ALPHABET]*Node
	end      bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	temp := &Trie{root: &Node{}}
	return temp
}

func (t *Trie) Insert(s string) {
	temp := t.root
	for i := 0; i < len(s); i++ {
		offset := s[i] - 'a'
		if temp.children[offset] == nil {
			temp.children[offset] = &Node{}
		}
		temp = temp.children[offset]
	}
	temp.end = true
}

func (t *Trie) Search(s string) bool {
	temp := t.root
	for i := 0; i < len(s); i++ {
		offset := s[i] - 'a'
		if temp.children[offset] == nil {
			return false
		}
		temp = temp.children[offset]
	}
	if temp.end == true {
		return true
	}
	return false
}
