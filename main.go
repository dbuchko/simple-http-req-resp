package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "net/http"
)

const (
      DEFAULT_PORT = "8080"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        t1 := time.Now()
        req_id := r.Header.Get("vcap_request_id")

        // Sleep for 1 second
        time.Sleep(50 * time.Millisecond)
        t2 := time.Now()

        nsec := t2.UnixNano() - t1.UnixNano()
        fmt.Println("App processing time for vcap_req_id=", req_id, " is ", nsec, "nanoseconds")

    default:
        fmt.Println("Sorry, only GET method is supported.")
    }
}

func main() {

    var port string
    if port = os.Getenv("PORT"); len(port) == 0 {
       fmt.Println("Warning, PORT not set. Defaulting to %+v", DEFAULT_PORT)
        port = DEFAULT_PORT
    }

    http.HandleFunc("/", handleGet)

    fmt.Println("Starting server for testing HTTP GET...")
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
