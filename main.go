package main

import (
    "log"
    "os"
    "time"
    "net/http"
)

const (
      DEFAULT_PORT = "8080"
)

var l = log.New(os.Stdout, "", 0)

func Log(msg string) {
        l.SetPrefix(time.Now().Format("2006-01-02 15:04:05.000") + " [INFO] ")
        l.Print(msg)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        Log("Received a GET request.")

    default:
        Log("Sorry, only GET method is supported.")
    }
}

func main() {

    var port string
    if port = os.Getenv("PORT"); len(port) == 0 {
       log.Printf("Warning, PORT not set. Defaulting to %+v", DEFAULT_PORT)
        port = DEFAULT_PORT
    }

    http.HandleFunc("/", handleGet)

    Log("Starting server for testing HTTP GET...")
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
