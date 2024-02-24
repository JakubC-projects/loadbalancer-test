package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	addr = ":80"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
