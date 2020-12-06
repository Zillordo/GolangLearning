package main

import (
	"./webportal"
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Webserver string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		log.Fatal(err)
	}

	err = webportal.RunWebPortal(config.Webserver)
	if err != nil {
		log.Fatal(err)
	}
}
