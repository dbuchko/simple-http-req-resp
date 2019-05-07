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
        fmt.Println("Received GET request for vcap_req_id=", req_id, " at ", t1.Format("2006-01-02T15:04:05.999999-07:00"))

        t2 := time.Now()
        fmt.Println("Processed GET request for vcap_req_id=", req_id, " at ", t2.Format("2006-01-02T15:04:05.999999-07:00"))

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
