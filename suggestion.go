package main

// suggestions given a prefix
// iterate over children nodes in trie in alphabetic order
// grab keys out of map, sort them, and iterate the sorted order of keys
// humanize package, values like long int and humanize the output
// 

import (
    "net/http"
    "fmt"
    "encoding/json"
    "log"
    "runtime"
    "strconv"
    "github.com/linneakw/suggestion/trie"
)

// TrieResponse represents a response from the suggestion route
type TrieResponse struct {
    Results *[]string `json:"Results"`
    Service string `json:"Service"`
}

var memstats = new(runtime.MemStats)

func getMemStats(w http.ResponseWriter, r *http.Request) {
    runtime.ReadMemStats(memstats)
    // fil out memstats with the current memory state
    // check golang.org/pkg/runtime/#MemStats
    
    allocStats := make(map[string]uint64)
    allocStats["alloc"] = memstats.Alloc
    allocStats["totalAlloc"] = memstats.TotalAlloc
    
    j, err := json.Marshal(allocStats)
    if nil != err {
        log.Println(err)
        w.WriteHeader(500)
        w.Write([]byte(err.Error()))
    } else {
        w.Header().Add("Content-Type", "application/json")
        w.Write(j)
        
    }
}


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
    
    // new instance by Name{}
    
    // Marshal usually refers to moving something over a network
    // also serializing
    // if nothing happens, j will be nil and err will be something
    // if somethign happens, j would exist and err would be nil
    j, err := json.Marshal(resp)
    
    if nil != err {
        log.Println(err)
        // specify http status code
        w.WriteHeader(598)
        w.Write([]byte("Results are still loading, stay tuned"))
        runtime.Gosched()
    } else {
        w.Header().Add("Content-Type", "application/json")
        w.Write(j)
        log.Println(resp.Results)
    }
    
    
    //w.Write([]byte(*resp.Results))
}

func main() {
    //http.HandleFunc("/", sayHello) 
    // change to handling a file server, serve static files that match 
    // the resource path out of the static folder
    // handle is what you use when the thing you pass to it has more capabilities
    http.Handle("/", http.FileServer(http.Dir("./static")))    
    
    // handle func expects second argument as go file
    http.HandleFunc("/api/v1/memstats", getMemStats)
    
    http.HandleFunc("/api/v1/suggestion", returnTrie)
    
    fmt.Println("Server listening on port 9000")
    // server is running, won't go past this line
    // otherwise, it would just stop the server
    http.ListenAndServe(":9000", nil)
    // if you want to listen on a network device, you would listen on that ip and then :
    
}