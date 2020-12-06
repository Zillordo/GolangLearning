package main

import (
	model "../../models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	data := &model.Animal{
		AnimalType: "second",
		Nickname:   "second",
		Zone:       3,
		Age:        44,
	}
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(data)
	if err != nil {
		log.Fatal(err)
	}

	//resp, err := http.Post("http://localhost:8080/api/dinos/add", "application/json", &b)
	//if err != nil || resp.StatusCode != 200 {
	//	log.Fatal(err)
	//}

	resp, err := http.Post("http://localhost:8080/api/dinos/edit/third", "application/json", &b)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal(resp.Status, err)
		return
	}
}
