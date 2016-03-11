package trie

import (
   "os" 
   "fmt"
   "bufio"
   "log"
   "io"
   "strings"
   "runtime"
   "sync"
)

// Trie type
type Trie struct {
    Root *TrieNode
    Words int
    mx sync.RWMutex
    SuggestionService string
    WG *sync.WaitGroup
}

// SearchTrie exports current instance of new trie
var SearchTrie = NewTrie()

// NewTrie constructor for the trie
func NewTrie() *Trie {
    // head node?
    dir := "./data/"
    NewTrie := &Trie{Root: NewTrieNode("")}
    //file := os.Args[1]
    if (len(os.Args) > 2) {
        file := string(os.Args[1])
        go NewTrie.ReadFile(dir + file);
        fmt.Println("Reading " + file)
        NewTrie.SuggestionService = "https://en.wikipedia.org/wiki/"
    } else {
        go NewTrie.ReadFile(dir + "wordsEn.txt")
        fmt.Println("Reading " + dir + "wordsEn.txt")
        NewTrie.SuggestionService = "http://www.dictionary.com/browse/"
    }
    
    return NewTrie
}

// ReadFile reads the file
func (trie *Trie) ReadFile(infilePath string) {
        //time.Sleep(5 * time.Second)
        // open the file
        infile, err := os.OpenFile(infilePath, os.O_RDONLY, 0644)
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
    var memstats = new(runtime.MemStats)
    runtime.ReadMemStats(memstats)
    // fill out memstats with the current memory state
    wg := sync.WaitGroup{}
    scanner := bufio.NewScanner(stream)
        scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        if (memstats.TotalAlloc >= 90000000) {
            wg.Wait()
        }   
        
        wg.Add(1)
        trie.AddEntry(scanner.Text(), &wg)
        
        
        // each line
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
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
func (trie *Trie) AddEntry(entry string, wg *sync.WaitGroup) {
    trie.mx.Lock()
    trie.Words++
    entry = strings.ToLower(entry)
    trie.Root.AddEntrie(entry)
    trie.mx.Unlock()
}

// AddEntrie takes a string and adds it to the trie
func (node *TrieNode) AddEntrie(entry string) *TrieNode { 
    if (len(entry) == 0) {
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
    
    child, err := node.children[first];
    
    if !err {
        child = NewTrieNode(first)
        node.children[first] = child
    }
    return child.AddEntrie(rest)
}


// FindEntries finds the matches of a given prefix
func (trie *Trie) FindEntries(prefix string, max uint8) *[]string {
    matches := make([]string, 0, max)
   
    node := trie.Root // start with the root
    count := 0
    // go through all of the prefix items
    for i := 0; i < len(prefix); i++ {
        // go over all of the node's children
        if len(node.children) > 0 {
            for key, child := range node.children {
                if (strings.ToLower(string(prefix[i])) == strings.ToLower(key)) {
                    if child != nil {
                        node = child
                        count++ 
                    }
                }
            }
        }
        // if the node hasn't been updated after ranging all of the current root's items

        
        // otherwise keep going!
    }
    if (count != len(prefix)) {
        return nil
    }
    // I know this isn't necessary but I'm trying everything to get rid of this error
    if node != nil {
        match(node, &matches, strings.ToLower(prefix), max)
    }
    
    
    return &matches
    
    // PHASE TWO
    // return complete items from that node
}

func match(node *TrieNode, matches *[]string, prefix string, max uint8) {
    if (node != nil && node.complete && len(*matches) < int(max)) {
        *matches = append(*matches, prefix)
    } 
    if (len(node.children) > 0 ) {
        for key, child := range node.children {
            newEntry := prefix + key
            if child != nil {
                match(child, matches, newEntry, max)
            }
            
        } 
    }
}