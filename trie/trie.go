package trie

import "fmt"

type Node struct {
	isWord bool
	next   map[string]*Node
}

type Trie struct {
	root *Node
	size int
}

func NewTrie() *Trie {
	t := &Trie{
		root: &Node{
			isWord: false,
			next:   make(map[string]*Node),
		},
	}
	return t
}

func (t *Trie) Add(word string) {
	cur := t.root
	for _, ch1 := range word {
		letter := fmt.Sprintf("%c", ch1)
		if cur.next[letter] == nil {
			// 不存在就添加
			cur.next[letter] = &Node{
				isWord: false,
				next:   make(map[string]*Node),
			}
		}
		// 将current指向下一级
		cur = cur.next[letter]
	}
	if cur.isWord == false {
		cur.isWord = true
		t.size++
	}
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for _, ch1 := range word {
		letter := fmt.Sprintf("%c", ch1)
		if cur.next[letter] == nil {
			return false
		}
		cur = cur.next[letter]
	}
	if cur.isWord {
		return true
	}
	return false
}

func (t *Trie) IsPrefix(word string) bool {
	cur := t.root
	for _, ch1 := range word {
		letter := fmt.Sprintf("%c", ch1)
		if cur.next[letter] == nil {
			return false
		}
		cur = cur.next[letter]
	}
	return true
}

func (t *Trie) showAll() {
	next := t.root.next["a"].next
	for key, value := range next {
		fmt.Printf("key : %s  value: %+v \n", key, value)
	}
}
