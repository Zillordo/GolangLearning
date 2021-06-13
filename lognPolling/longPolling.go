package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var messages = make(chan string)

// standard http request
func pollResponse(w http.ResponseWriter, _ *http.Request) {
	timeout := make(chan bool)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// lambda function that runs as concurrent function
	go func() {
		//tweak this number for different timeout
		time.Sleep(30e9)
		timeout <- true
	}()

	//this select waits for whatever comes first (message, timeout)
	select {
	case msg := <-messages:
		_, _ = io.WriteString(w, msg)
	case <-timeout:
		return
	}
}

// this function generates random data in random time as a concurrent task
func generateRandomEvents() {
	events := []string{
		"Notification: error",
		"Notification: notification",
		"Some new notification",
		"Hello from server",
		"Some random notification",
		"This was done by long pooling",
	}
	for true {
		//tweak this number to change the range of time
		time.Sleep(time.Duration(rand.Intn(5)) * time.Minute)
		messages <- events[rand.Intn(len(events))]
	}
}

func main() {
	// http handleFunc will add new handler to to golangs multiplexer
	http.HandleFunc("/events", pollResponse)
	go generateRandomEvents()

	fmt.Println("server is listening on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
