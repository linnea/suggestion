package trie

import (
   "testing"
   "os"
)
// TestTrie tests the trie
func TestTrie(t *testing.T) {
    
    
    trie := NewTrie()
    argFile := string(os.Args[1])
    
    if (len(argFile) > 1) {
        if (trie.SuggestionService != "https://en.wikipedia.org/wiki/") {
            t.Error("Wrong suggestion service found")
        } else {
            t.Log("The suggestion service is correct")
        }
    } else {
        if (trie.SuggestionService != "http://www.dictionary.com/browse/") {
            t.Error("Wrong suggestion service found")
        } else {
            t.Log("The suggestion service is correct")
        }
    }
    
    
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
    matches := *trie.FindEntries("aah", max)
    if (len(matches) != 4) {
        t.Error("Did not find appropriate amount of matches for 'aah'")
    }
    
    matches = *trie.FindEntries("a", max) 
    if (len(matches) > int(max)) {
        t.Error("Max exceeded, maximum is not valid")
    } else {
        t.Log("Max enabled and verified")
    }
}
// returning capitalized matches

