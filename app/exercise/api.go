package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bryanl/apptracing/internal/platform/db"
	"github.com/gorilla/mux"
)

func initHTTP() http.Handler {
	opts := "user=postgres dbname=velocity2017 sslmode=disable"
	ds := db.Init("postgres", opts)

	r := mux.NewRouter()
	r.HandleFunc("/people", peopleHandler(ds)).Methods("GET")
	r.HandleFunc("/people/{id}", personHandler(ds)).Methods("GET")

	return r
}

type peopleResponse struct {
	People []person `json:"people,omitempty"`
}

type personResponse struct {
	Person person `json:"person,omitempty"`
}

type errorResponse struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func peopleHandler(ds *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		page := r.URL.Query().Get("page")
		if page == "" {
			page = "1"
		}

		perPage := r.URL.Query().Get("per_page")
		if perPage == "" {
			perPage = "50"
		}

		pageI, _ := strconv.Atoi(page)
		perPageI, _ := strconv.Atoi(perPage)

		w.Header().Set("Content-Type", "application/json")

		people, err := peopleListService(ctx, ds, pageI, perPageI)
		if err != nil {
			resp := errorResponse{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			}

			w.WriteHeader(resp.Code)
			_ = json.NewEncoder(w).Encode(&resp)

			return
		}

		w.WriteHeader(http.StatusOK)

		resp := peopleResponse{People: people}
		_ = json.NewEncoder(w).Encode(&resp)
	}
}

func personHandler(ds *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var p person
		if err := ds.GetContext(r.Context(), &p, queryPerson, vars["id"]); err != nil {
			http.Error(w, fmt.Sprintf("unable to load person: %v", err), http.StatusInternalServerError)
		}

		resp := personResponse{Person: p}
		json.NewEncoder(w).Encode(&resp)
	}
}
