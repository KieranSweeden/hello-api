package main

import (
	"log"
	"net/http"

	"github.com/kieransweeden/hello/handlers/rest"
)

type Response struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
