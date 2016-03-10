package trie 

// CASE SENSITIVE ATTEMPT
/*
func (trie *Trie) FindEntries(prefix string, max uint8) []string {
    /* CASE SENSITIVE ATTEMPT
    matches := make([]string, max)
    // for comparisons
    lowerPrefix := strings.ToLower(prefix)
    // go through the trie
    
    // modifies nodes to be pointing at
    // the viable prefix nodes
    nodes := make([]TrieNode, max)
    nodes = append(nodes, *trie.Root)
    for i := 0; i < len(prefix); i++ {
        searchTrie(nodes, string(lowerPrefix[i]))
    }
    
    // now go through those prefix nodes!
    // and find the matches
    
    
    // for the length of the prefix string
        // if key.tolowercase = firstletter.tolowercase
            // continue
            // pass it findmatch(prefix[nextIndex])
        // if .tolowercase=
    // DO SOMETHING WITH NODES
    fmt.Println("Remaining nodes searchable after prefix" )
    for i := 0; i < len(nodes); i++ {
        fmt.Println(nodes[i])
    }
    return match(prefix, nodes, matches)
}*/
/* Again attempting case sensitive
// searches the trie given a starting node and key
func searchTrie(nodes []TrieNode, startKey string) {
    
    /// MAYBE NEED A POINTER TO TRIE NODE ARRAY
    // for all viable nodes
    for i := 0; i < len(nodes); i++ {
        // and for all of that nodes children
        for key, node := range nodes[i].children {
            // go through the current level children
            if (startKey == strings.ToLower(key)) {
                nodes = append(nodes, *node)
            }
        }
        fmt.Println("removing", nodes[i])
        
        // then remove that node since it's been searched
        nodes = append(nodes[:i], nodes[i+1:]...)
        
    } 
}

func match(prefix string, nodes []TrieNode, matches []string) []string {
    // while the last item in the array isn't filled
    for (matches[len(matches) - 1] != "" && len(nodes) != 0) {
        // if you find a complete entry, add it
        node := nodes[0]
        if (node.complete) {
            matches = append(matches, prefix)
            fmt.Println("appending to matches", prefix)
        } 
        if (len(node.children) > 0 ) {
            for key, child := range node.children {
                newEntry := prefix + key
                //fmt.Print(string(entry[:len(entry) - 1]), ": ")
                nodes = append(nodes, *child)
                
                match(newEntry, nodes, matches)
            } 
        }
        
        nodes = append(nodes[:0], nodes[1:]...)
    }
    return matches
}*/

// another function that accepts a file path, opens it as an io.Reader stream, and calls the ffunc
// passed in a file path? 

