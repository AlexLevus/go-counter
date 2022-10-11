package main

import (
    "fmt"
    "net/http"
)

var counter int = 0;

func showCounter(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Счетчик равен %v\n", counter)
}

func increment(w http.ResponseWriter, req *http.Request) {
	counter = counter + 1;
    fmt.Fprintf(w, "Счетчик равен %v\n", counter)
}

func showName(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "<h3> Hello, Александр Левусь</h3>")
}

func main() {
    http.HandleFunc("/", showCounter)

    http.HandleFunc("/stat", increment)
    http.HandleFunc("/about", showName)

	fmt.Printf("%v\n", "Listen 8080 port\n")
    http.ListenAndServe(":8080", nil)
}