package main

import "fmt"

type trie struct {
	root *node
}

// Add adds a word to the trie
func (t *trie) Add(word string) {
	t.root.add([]rune(word))
}

// SearchWord checks if a specific word exists in the trie
func (t *trie) SearchWord(word string) bool {
	lastNode, ok := t.search(word)
	return ok && lastNode.endOfWord
}

// search searches for string. Returns last node and true if found
func (t *trie) search(str string) (lastNode *node, ok bool) {
	crawl := t.root
	for _, c := range str {
		crawl, ok = crawl.getChild(c)
		if !ok {
			return
		}
	}
	lastNode = crawl
	ok = true
	return
}

// StartsWith returns all the words that starts with the given string
func (t *trie) StartsWith(str string) (words []string) {
	crawl, ok := t.search(str)
	if !ok {
		return
	}
	words = getStringFromNode(crawl, str)
	return
}

func getStringFromNode(n *node, prefix string) (result []string) {
	for _, child := range n.children {
		result = append(result, getStringFromNode(child, fmt.Sprintf("%s%c", prefix, child.char))...)
	}
	if n.endOfWord {
		result = append(result, prefix)
	}
	return result
}

type node struct {
	char      rune
	children  map[rune]*node
	endOfWord bool
}

func (n *node) add(chars []rune) {
	if len(chars) == 0 {
		n.endOfWord = true
		return
	}
	c := chars[0]
	nodeNew, ok := n.children[c]
	if !ok {
		nodeNew = newNode(c)
		n.children[c] = nodeNew
	}
	nodeNew.add(chars[1:])
}

func (n *node) getChild(c rune) (child *node, ok bool) {
	child, ok = n.children[c]
	return
}
func newTrie() *trie {
	return &trie{
		root: newNode('*'),
	}
}

func newNode(c rune) *node {
	return &node{
		char:     c,
		children: make(map[rune]*node),
	}
}
