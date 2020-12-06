package api

import (
	"../../database"
	"github.com/gorilla/mux"
	"net/http"
)

func RunApi(endpoint string, db database.DinoDBHandler) error {
	r := mux.NewRouter()
	RunAPIOnRouter(r, db)
	return http.ListenAndServe(endpoint, r)
}

func RunAPIOnRouter(r *mux.Router, db database.DinoDBHandler) {
	handler := newDinoRESTAPIHandler(db)

	apiRouter := r.PathPrefix("/api/dinos").Subrouter()

	apiRouter.Methods(http.MethodGet).Path("/{SearchCriteria}/{search}").HandlerFunc(handler.searchHandler)
	apiRouter.Methods(http.MethodPost).PathPrefix("/{Operation}").HandlerFunc(handler.editsHandler)
}
