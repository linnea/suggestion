package trie

import (
   "testing"
   //"fmt"
)
// TestTrie tests the trie
func TestTrie(t *testing.T) {
    trie := NewTrie("")
    trie.AddEntry("Yay")
    trie.AddEntry("Yaah")
    
    trie.PrintTree()
}
