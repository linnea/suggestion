package trie

import (
   "os" 
   "bufio"
   "log"
   "io"
   "github.com/linneakw/suggestion/node"
)

// Trie type
type Trie struct {
    root *TrieNode
}

// NewTrie constructor for the trie
func NewTrie(file string) *Trie {
    // head node?
    
    if (len(file) > 0) {
        ReadFile(file);
    } else {
        ReadFile("../data/wordsEn.txt")
    }
    
    return &Trie{root: NewTrieNode()}
}




// LoadTrie loads the trie
func LoadTrie(stream io.Reader) {
    scanner := bufio.NewScanner(stream)
        scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        //this.AddEntry(scanner.Text())
        // each line
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }/*
    strings.NewReader(stream)
    
    s := "Hello World!"
    
    for idx, r := range s {
        fmt.Println(idx, r)
    }*/
}

// printHelper helps print out the trie
func PrintTrie(node *Trie,) {
    fmt.Println(printHelper("", root))
}

// PrintTree prints out the tree
// Worked with Conrad to understand implementation style
func (trie *Trie) printHelper(entry string, node *TrieNode) {
    
    //fmt.Println()
    count := 0
    if (node.complete) {
        fmt.Println(entry)
    } else {
        for key, child := range trie.children {
            string := string(key[0]))
            if (trie.complete) {
                return string
            } else {
                fmt.Print(" ")
            }
            count++
            string += child.printHelper(string, child)
            return string
        }
        for i := 0; i < count; i++ {
            fmt.Print("| ")
        }
    }
}



// AddEntry takes a string and adds it to the trie
func (trie *Trie) AddEntry(entry string) *Trie { 
    if len(entry) == 0 {
        trie.complete = true;
        return trie
    }
    /*
    // if numwords < 50
    if (trie.NumWords < 50) {
        for index, runeValue := range string {
            trie.Children[runeValue]
        }
    }
        // go to head node
        // try to find first letter in NODES map
        // if doesn't match string (or not first letter) {
            // add whole word to map
        //} else if numwords ==50 {
            // reformat trie node's children
        //}*/
    
    first := string(entry[0])
    rest := string(entry[1:])
    
    child, ok := trie.children[first];
    
    if !ok {
        child = NewTrie(first)
        trie.children[first] = child
    }
    return child.AddEntry(rest)
}
// FindEntries
//func (trie *Trie) FindEntries(prefix string, max uint8) []string {}

// another function that accepts a file path, opens it as an io.Reader stream, and calls the ffunc
// passed in a file path? 

