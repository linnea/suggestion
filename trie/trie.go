package trie

import (
   "os" 
   "fmt"
   "bufio"
   "log"
   "io"
   "strings"
)

// Trie type
type Trie struct {
    Root *TrieNode
    Words int
}

// NewTrie constructor for the trie
func NewTrie() *Trie {
    // head node?
    dir := ".././data/"
    NewTrie := &Trie{Root: NewTrieNode("")}
    //file := os.Args[1]
    if (len(os.Args) > 1) {
        file := string(os.Args[1])
        NewTrie.ReadFile(dir + file);
        fmt.Println("Reading " + file)
    } else {
        NewTrie.ReadFile(dir + "text.txt")
    }
    
    return NewTrie
}

// ReadFile reads the file
func (trie *Trie) ReadFile(infilePath string) {
        // open the file
        infile, err := os.Open(infilePath)
        if err != nil {
            log.Fatal(err)
        }
        defer infile.Close()

        trie.LoadTrie(infile)
        
        //creates and loads a trie from any io.Reader stream
        //var stream := io.Reader
        //loadTrie(stream.Read(n int, err error)
    }


// LoadTrie loads the trie
func (trie *Trie) LoadTrie(stream io.Reader) {
    scanner := bufio.NewScanner(stream)
        scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        trie.AddEntry(scanner.Text())
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

// PrintHelper prints out the tree
// Worked with Conrad to understand implementation style
func printHelper(entry string, node *TrieNode) {
    //fmt.Println()
    if (node.complete) {
        fmt.Println(entry)
    } 
    if (len(node.children) > 0 ) {
        for key, child := range node.children {
            newEntry := entry + key
            //fmt.Print(string(entry[:len(entry) - 1]), ": ")
            printHelper(newEntry, child)
        } 
    }
}

// PrintTrie printHelper helps print out the trie
func (trie *Trie) PrintTrie() {
    printHelper("", trie.Root)
}

// AddEntry it into the trie
// lower case for simplicity
func (trie *Trie) AddEntry(entry string) {
    trie.Words++
    strings.ToLower(entry)
    trie.Root.AddEntrie(entry)
}

// AddEntrie takes a string and adds it to the trie
func (node *TrieNode) AddEntrie(entry string) *TrieNode { 
    if len(entry) == 0 {
        node.complete = true;
        return node
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
    
    child, ok := node.children[first];
    
    if !ok {
        child = NewTrieNode(first)
        node.children[first] = child
    }
    return child.AddEntrie(rest)
}


// FindEntries finds the matches of a given prefix
func (trie *Trie) FindEntries(prefix string, max uint8) *[]string {
    matches := make([]string, 0, max)
    prefix = strings.ToLower(prefix)
    node := trie.Root // start with the root
    updated := false
    // go through all of the prefix items
    for i := 0; i < len(prefix); i++ {
        // go over all of the node's children
        for key, child := range node.children {
            if (string(prefix[i]) == key) {
                node = child
                updated = true
            }
        }
        // if the node hasn't been updated after ranging all of the current root's items
        if (!updated) {
            return nil
        }
        
        // otherwise keep going!
    }
    
    match(node, &matches, prefix, max)
    
    return &matches
    
    // PHASE TWO
    // return complete items from that node
}

func match(node *TrieNode, matches *[]string, prefix string, max uint8) {
    if (node.complete && uint8(len(*matches)) < max) {
        *matches = append(*matches, prefix)
    } 
    if (len(node.children) > 0 ) {
        for key, child := range node.children {
            newEntry := prefix + key
           
            match(child, matches, newEntry, max)
        } 
    }
}