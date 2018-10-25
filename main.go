package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Status struct
type Status struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	Creator     string `json:"creator"`
	PublishDate string `json:"publishDate"`
}

// Account / Tag struct
type Account struct {
	Id         string `json:"id"`
	Uid        string `json:"uid"`
	Name       string `json:"name"`
	TagName    string `json:"tagName"`
	Permission string `json:"permission"`
}

// Access struct
type Access struct {
	id         int
	account    Account
	accessDate time.Time
}

// Get current Status
func getStatuses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(GetData("statusLog"))
	//fmt.Println(json.NewEncoder(w).Encode())
}

// Update current Status
func updateStatus(w http.ResponseWriter, r *http.Request) {

}

// Check if UID has access
func hasAccess(w http.ResponseWriter, r *http.Request) {

}

// Open the door
func open(w http.ResponseWriter, r *http.Request) {

}

// Log to DB
func logToDB(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/getStatuses", getStatuses).Methods("GET")
	r.HandleFunc("/api/addStatus", updateStatus).Methods("POST")
	r.HandleFunc("/api/hasAccess/{uid}", hasAccess).Methods("GET")
	r.HandleFunc("/api/open", open).Methods("GET")
	r.HandleFunc("/api/log", logToDB).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
