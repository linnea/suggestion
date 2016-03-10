package main

// suggestions given a prefix
// iterate over children nodes in trie in alphabetic order
// grab keys out of map, sort them, and iterate the sorted order of keys
// humanize package, values like long int and humanize the output
// 

import (
    "net/http"
    "fmt"
    "time"
    "encoding/json"
    "log"
    "runtime"
)

// HelloResponse represents a response from the hello route
type HelloResponse struct {
    Name string `json:"name"`
    Message string `json:"message"`
    GeneratedAt time.Time `json:"generatedAt"`
    foo int
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

func sayHello(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Path[len("/api/v1/hello/"):]
    resp := HelloResponse{
        Name: name, 
        Message: "Hello " + name, 
        GeneratedAt: time.Now()}
    // new instance by Name{}
    
    // Marshal usually refers to moving something over a network
    // also serializing
    // if nothing happens, j will be nil and err will be something
    // if somethign happens, j would exist and err would be nil
    j, err := json.Marshal(resp)
    
    if nil != err {
        log.Println(err)
        // specify http status code
        w.WriteHeader(500)
    } else {
        w.Header().Add("Content-Type", "application/json")
        w.Write(j)
    }
    
    
    w.Write([]byte("Hello " + name))
}

func main() {
    //http.HandleFunc("/", sayHello) 
    // change to handling a file server, serve static files that match 
    // the resource path out of the static folder
    // handle is what you use when the thing you pass to it has more capabilities
    http.Handle("/", http.FileServer(http.Dir("./static")))    
    
    // handle func expects second argument as go file
    http.HandleFunc("/api/v1/hello/", sayHello)
    http.HandleFunc("/api/v1/memstats", getMemStats)
    
    fmt.Println("Server listening on port 9000")
    // server is running, won't go past this line
    // otherwise, it would just stop the server
    http.ListenAndServe(":9000", nil)
    // if you want to listen on a network device, you would listen on that ip and then :
}