package main

import (
 "fmt"
 "log"
 "net/http"
)

func main() {
 http.HandleFunc("/", handler)
 fmt.Println("> :8080");
 log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprint(w, "Hello, World!")
}