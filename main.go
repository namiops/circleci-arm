package main

import (
        "fmt"
        "log"
        "net/http"
)

func main() {
        defaultHandler := func(write http.ResponseWriter, req *http.Request) {
                fmt.Fprintf(write, "Golang Arm Docker listen on port: 5050  request method %s, URL: %s\n", req.Method, req.URL)
        }

        http.HandleFunc("/", defaultHandler)
        log.Fatal(http.ListenAndServe(":5050", nil))
}
