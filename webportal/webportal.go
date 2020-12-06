package webportal

import (
	"fmt"
	"net/http"
)

func RunWebPortal(addr string) error {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(addr, nil)
	return err
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome to the web portal %s", r.RemoteAddr)
}
