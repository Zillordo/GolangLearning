package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", loo)
	log.Fatal(http.ListenAndServeTLS(":10443", "webdevtoolkit/key/cert.pem", "webdevtoolkit/key/key.pem", nil))
}

func loo(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("This is an example server.\n"))
}
