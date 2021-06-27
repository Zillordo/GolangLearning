package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func index(w http.ResponseWriter, _ *http.Request) {
	_, _ = io.WriteString(w, "hello from a docker container 2")
}
