package trie 


// TrieNode for Trie
type TrieNode struct {
    children map[string]*TrieNode
    complete bool
}

// NewTrieNode acts as node constructor
func NewTrieNode(value string) *TrieNode {
    return &TrieNode{children: map[string]*TrieNode{}}
}