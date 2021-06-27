package app

import (
	"io"
	"log"
	"net/http"

	"../controllers"
)

func StartApplication() {
	http.HandleFunc("/", index)
	http.HandleFunc("/users", controllers.GetUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, _ *http.Request) {
	_, _ = io.WriteString(w, "hello form index")
}
