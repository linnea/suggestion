package main

import (
    "net/http"
    "fmt"
    "encoding/json"
    "log"
    "strconv"
    "flag"
    "github.com/linneakw/challenges-linneakw/suggestion/trie"
)

// TrieResponse represents a response from the suggestion route
type TrieResponse struct {
    Results *[]string `json:"Results"`
    Service string `json:"Service"`
}

 // searches the trie and returns the results of search query
 func returnTrie(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("q")
    max, err := strconv.Atoi(r.URL.Query().Get("max"))
    if err != nil {
        // if max isn't within acceptable range
        if (max > 50 || max < 10) {
            max = 20
        }
    }
    
    searchTrie := trie.SearchTrie
    resp := TrieResponse {
            Results: searchTrie.FindEntries(query, uint8(max)),
            Service: searchTrie.SuggestionService} 
    
    // new json response
    
    // marshal the response
    j, err := json.Marshal(resp)
    fmt.Println(resp.Results)
    if (nil != err) {
        log.Println(err)
        // specify http status code
        w.WriteHeader(598)
        w.Write([]byte("Results are still loading, stay tuned"))
       
    } else {
        w.Header().Add("Content-Type", "application/json")
        w.Write(j)
        log.Println(resp.Results)
    }
    
    
    //w.Write([]byte(*resp.Results))
}

func main() {
    // parse the command line arguments
    flag.Parse()
    
    // handle serving files
    http.Handle("/", http.FileServer(http.Dir("./static")))
    
    // expose api route
    http.HandleFunc("/api/v1/suggestion", returnTrie)
    
    // listen on port 9000
    fmt.Println("Server listening on port 9000")
    http.ListenAndServe(":9000", nil)
    
}