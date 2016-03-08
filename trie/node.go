package trie 

import (
    
)

// TrieNode for Trie
type TrieNode struct {
    children map[string]*TrieNode
    complete bool
}

// NewTrieNode acts as node constructor
func NewTrieNode() *node {
    return &node{children: map[string]*TrieNode{}}
}