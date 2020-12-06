package api

import (
	"../../database"
	model "../../models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type DinoRESTAPIHandler struct {
	dbHandler database.DinoDBHandler
}

func newDinoRESTAPIHandler(db database.DinoDBHandler) *DinoRESTAPIHandler {
	return &DinoRESTAPIHandler{dbHandler: db}
}

func (h *DinoRESTAPIHandler) searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `No search criteria found, zou can either search by nickname via /api/dinos/nickname/rex, or
								to search by type via /api/dinos/type/velociraptor`)
		return
	}

	searchKey, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `No search criteria found, zou can either search by nickname via /api/dinos/nickname/rex, or
								to search by type via /api/dinos/type/velociraptor`)
	}

	var animal model.Animal
	var animals []model.Animal
	var err error
	switch strings.ToLower(criteria) {
	case "nickname":
		animal, err = h.dbHandler.GetDynoByNickname(searchKey)
	case "type":
		animals, err = h.dbHandler.GetDynosByType(searchKey)
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "Error occured while querying animas %v", err)
		return
	}
	if len(animals) > 0 {
		_ = json.NewEncoder(w).Encode(animals)
		return
	}
	_ = json.NewEncoder(w).Encode(animal)
}

func (h *DinoRESTAPIHandler) editsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operation, ok := vars["Operation"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, "Operation was not provided")
		return
	}
	var animal model.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "Could not decode the request bodz to json %v", err)
		return
	}

	switch strings.ToLower(operation) {
	case "add":
		err = h.dbHandler.AddAnimal(animal)
	case "edit":
		nickname := r.RequestURI[len("/api/dinos/edit/"):]
		log.Println("edit requested for nickname", nickname)
		err = h.dbHandler.UpdateAnimal(animal, nickname)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "Error occured while processing request %v", err)
	}
}

func ToJSON(data interface{}) (*bytes.Buffer, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)

	return b, nil
}

func Decode(data []byte, t interface{}) error {
	return json.Unmarshal(data, t)
}
