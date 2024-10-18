package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func Handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
func counter(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "counter URL.Path = %q\n", req.URL.Path)
}
