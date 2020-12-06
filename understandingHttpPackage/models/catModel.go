package models

import (
	"fmt"
	"net/http"
)

type Cat struct {
	Name string
	Age  int
}

func (c *Cat) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.get(w)
	default:
		methodNotAllowed(w, r)
	}
}

func (c *Cat) get(w http.ResponseWriter) {
	c.Name = "Leo"
	c.Age = 2
	_, _ = fmt.Fprintf(w, "This is cat named: %s, and he is %d years old.", c.Name, c.Age)
}
