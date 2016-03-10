package trie

import (
   "testing"
   "fmt"
)
// TestTrie tests the trie
func TestTrie(t *testing.T) {
    trie := NewTrie()
    trie.AddEntry("Yay")
    trie.AddEntry("Yaah")
    // check that adding successfully by # words
    trie.PrintTrie();
    fmt.Println(trie.Words, "\n")
}

func TestMatch(t *testing.T) {
    trie := NewTrie()
    trie.AddEntry("Wow")
    trie.AddEntry("Dude")
    trie.AddEntry("woman")
    trie.AddEntry("Women")
    trie.AddEntry("Wowzers")
    // returning right amount of matches
    var max uint8 = 4
    matches := *trie.FindEntries("wo", max)
    matches2 := *trie.FindEntries("w", 2) 
    fmt.Println("4 matches")
    for i := 0; i < len(matches); i++ {
        fmt.Println(matches[i])
        
    }
    
    fmt.Println("\n2 matches")
    for i := 0; i < len(matches2); i++ {
        fmt.Println(matches2[i])
        
    }
    
    
    
    
}

// returning capitalized matches

