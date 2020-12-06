package router

import (
	"../models"
	"fmt"
	"net/http"
)

type Route struct {
	Url     string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var dog models.Dog
var cat models.Cat

var Routes = []Route{
	{
		Url:     "/",
		Handler: defaultHandler,
	},
	{
		Url:     "/dog",
		Handler: dog.Handler,
	},
	{
		Url:     "/cat",
		Handler: cat.Handler,
	},
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/" {
			defaultGet(w, r)
		} else {
			_, _ = fmt.Fprint(w, "404 not found")
		}
	}
}

func defaultGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "understandingHttpPackage/assets/page.html")
}
