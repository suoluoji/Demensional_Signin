package main

import (
    "net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello"))
}

func main() {
    http.HandleFunc("/hello", SayHello)
    err := http.ListenAndServe(":8001", nil)
    if err != nil {
        panic(err)
    }
}
