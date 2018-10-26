package main

import (
	"accessControl_v3/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/getStatuses", api.GetStatuses).Methods("GET")
	r.HandleFunc("/api/getStatus", api.GetStatus).Methods("GET")
	r.HandleFunc("/api/updateStatus", api.UpdateStatus).Methods("POST")
	r.HandleFunc("/api/hasAccess", api.HasAccess).Methods("POST")
	r.HandleFunc("/api/logAccess", api.LogAccess).Methods("POST")
	r.HandleFunc("/api/open", api.Open).Methods("GET")

	// Start server
	go log.Fatal(http.ListenAndServe(":8000", r))
}
