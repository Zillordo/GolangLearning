package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", somFunc)
	http.Handle("/favicon.icon", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func somFunc(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
