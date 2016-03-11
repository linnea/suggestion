package trie

import (
   "testing"
   "flag"
   "fmt"
)
// TestTrie tests the trie
func TestTrie(t *testing.T) {
    flag.Parse()
    
    trie := SearchTrie
    
    // test the amount of entries added
    words := trie.Words
    trie.AddEntry("yodel", trie.WG)
    trie.AddEntry("doodah", trie.WG)
    if (trie.Words != words + 2) {
        t.Error("Entering new entries either didn't increment amount of entries, or failed")
    } else {
        t.Log("Adding entries successful")
    }
    
    // Test that 
    // returning right amount of matches
    var max uint8 = 6
    matches := trie.FindEntries("aah", max)
    fmt.Println(matches)
    if (len(matches) != 4) {
        t.Error("Did not find appropriate amount of matches for 'aah', found ", len(matches))
    }
    
    matches = trie.FindEntries("a", max) 
    if (len(matches) > int(max)) {
        t.Error("Max exceeded, maximum is not valid")
    } else {
        t.Log("Max enabled and verified")
    }
}
// returning capitalized matches

