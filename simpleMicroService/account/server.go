package account

import (
	"context"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPServer(_ context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httpTransport.NewServer(endpoints.CreateUser, decodeUserReq, encodeResponse))
	r.Methods("GET").Path("/user/{id}").Handler(httpTransport.NewServer(endpoints.GetUser, decodeEmailReq, encodeResponse))
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
