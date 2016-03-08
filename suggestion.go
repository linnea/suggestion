package main

import (
    "github.com/linneakw/suggestion/trie"
    "os"
)

func func main() {
    // ReadFile is
    func ReadFile(infilePath string) {
        // open the file
        infile, err := os.Open(infilePath)
        if err != nil {
            log.Fatal(err)
        }
        defer infile.Close()

        Trie.LoadTrie(infile)
        
        //creates and loads a trie from any io.Reader stream
        //var stream := io.Reader
        //loadTrie(stream.Read(n int, err error)
    }
}