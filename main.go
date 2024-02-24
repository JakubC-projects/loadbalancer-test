package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	addr     = ":443"
	certFile = "/etc/letsencrypt/live/xn-5r8h.pl/fullchain.pem"
	keyFile  = "/etc/letsencrypt/live/xn-5r8h.pl/privkey.pem"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	if err := http.ListenAndServeTLS(addr, certFile, keyFile, mux); err != nil {
		log.Fatal(err)
	}
}
