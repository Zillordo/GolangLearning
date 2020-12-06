package models

import (
	"fmt"
	"net/http"
)

type Dog struct {
	Name string
	Age  int
}

func (d *Dog) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		d.get(w)
	default:
		methodNotAllowed(w, r)
	}
}

func (d *Dog) get(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	d.Name = "Aika"
	d.Age = 5
	_, _ = fmt.Fprintf(w, "This is dog named: %s, and he is %d years old.", d.Name, d.Age)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Method %s not allowed", r.Method)
}
