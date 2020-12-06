package main

import (
	"./router"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("understandingHttpPackage/assets"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	for _, route := range router.Routes {
		http.HandleFunc(route.Url, route.Handler)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
