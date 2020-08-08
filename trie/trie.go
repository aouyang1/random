package main

import (
//"fmt"
)

type node struct {
	children map[byte]*node
	isWord   bool
}

func newNode() *node {
	return &node{children: make(map[byte]*node)}
}

// insert stores the word byte slice into the trie and returns if it's a new word
func (n *node) insert(w []byte) bool {
	if len(w) == 0 {
		// already inserted
		if n.isWord {
			return false
		}

		// new word
		n.isWord = true
		return true
	}
	cn, exists := n.children[w[0]]
	if !exists {
		cn = newNode()
		n.children[w[0]] = cn
	}
	return false || cn.insert(w[1:])
}

// remove looks for a word and removes the appropriate nodes if there are no
// children remaining. This returns two booleans, one on telling the parent node
// to remove itself, and two if the word will be deleted
func (n *node) remove(w []byte) (bool, bool) {
	// found the word
	if len(w) == 0 {
		if n.isWord {
			// found the word
			n.isWord = false
			if len(n.children) == 0 {
				return true, true
			}
			return false, true
		}
		return false, false
	}
	// see if children has the start of the word
	cn, exists := n.children[w[0]]
	if !exists {
		// didn't find the word
		return false, false
	}
	removeChild, removing := cn.remove(w[1:])
	// this means the returned child has no other children and can be removed
	if removeChild {
		delete(n.children, w[0])
	}

	// check if current node still has children and if it isn't a word
	// then this node can also be removed
	return len(n.children) == 0 && !n.isWord, removing
}

// findPrefix searches for the node after finding the prefix in the trie and returns nil if it cannot be found
func (n *node) findPrefix(p []byte) *node {
	if len(p) == 0 {
		return n
	}
	cn, exists := n.children[p[0]]
	if !exists {
		return nil
	}
	return cn.findPrefix(p[1:])
}

// search finds all possible words from this node
func (n *node) search(p []byte) []string {
	var words []string

	// arrived at a node that indicates word so save it
	if n.isWord {
		words = append(words, string(p))
	}

	// check if we can go further down
	for k, cn := range n.children {
		words = append(words, cn.search(append(p, k))...)
	}
	return words
}

type Trie struct {
	root     *node
	numWords int
}

func NewTrie() *Trie {
	return &Trie{root: newNode()}
}

// NumWords returns the number of words stored in the Trie
func (t *Trie) NumWords() int {
	return t.numWords
}

// Insert stores a word in the trie data structure. Returns true if the word
// is not currently present in the trie and was added successfully.
func (t *Trie) Insert(w string) bool {
	if len(w) == 0 {
		return false
	}
	wb := []byte(w)
	if t.root.insert(wb) {
		t.numWords++
		return true
	}
	return false
}

// Remove looks for the given string and removes it from the trie if present.
// Returns true if the word is found in the trie and successfully removed.
func (t *Trie) Remove(w string) bool {
	wb := []byte(w)
	_, removed := t.root.remove(wb)
	if removed {
		t.numWords--
		return true
	}
	return false
}

// Search finds all possible words with a given prefix
func (t *Trie) Search(p string) []string {
	pb := []byte(p)
	n := t.root.findPrefix(pb)
	if n == nil {
		return nil
	}
	res := n.search([]byte(""))
	for i, r := range res {
		res[i] = p + r
	}
	return res
}
