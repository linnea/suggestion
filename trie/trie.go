package trie

import (
   "os" 
   "fmt"
   "bufio"
   "log"
   "io"
)

// Trie type
type Trie struct {
    NumWords int
    Children map[string]*Trie
    complete bool
}

// NewTrie constructor for the trie
func NewTrie(file string) *Trie {
    // head node?
    
    if (len(file) > 0) {
        ReadFile(file);
    } else {
        ReadFile("../data/wordsEn.txt")
    }
    
    return &Trie{Children: map[string]*Trie{}}
}

// AddEntry takes a string and adds it to the trie
func (trie *Trie) AddEntry(entry string) { 
    fmt.Println("Entry is", entry)
    // if numwords < 50
        // go to head node
        // try to find first letter in NODES map
        // if doesn't match string (or not first letter) {
            // add whole word to map
        //} else {
            // 
        //}
        /*
    for index, runeValue := range string {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }*/
}
// FindEntries
//func (trie *Trie) FindEntries(prefix string, max uint8) []string {}

// ReadFile is
func ReadFile(infilePath string) {
    // open the file
    infile, err := os.Open(infilePath)
    if err != nil {
        log.Fatal(err)
    }
    defer infile.Close()

    loadTrie(infile)
    
    //creates and loads a trie from any io.Reader stream
    //var stream := io.Reader
    //loadTrie(stream.Read(n int, err error)
}

// LoadTrie loads the trie
func loadTrie(stream io.Reader) {
    scanner := bufio.NewScanner(stream)
        scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
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

// another function that accepts a file path, opens it as an io.Reader stream, and calls the ffunc
// passed in a file path? 

