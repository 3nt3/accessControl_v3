package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Status struct
type Status struct {
	ID          int    `json:"id"`
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
	account    string
	accessDate time.Time
}

// Get current Status
func getStatuses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetData("statusLog"))
}

// Update current Status
func updateStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newStatus Status
	_ = json.NewDecoder(r.Body).Decode(&newStatus)

	data := []interface{}{newStatus.Status, newStatus.Creator}
	InsertData("statusLog", data)
	json.NewEncoder(w).Encode(GetData("statusLog")[len(GetData("statusLog"))-1])
}

// Check if UID has access
func hasAccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonData map[string]string
	_ = json.NewDecoder(r.Body).Decode(&jsonData)

	uid := jsonData["uid"]
	var uidHasAccess bool

	for _, item := range GetData("accounts") {
		var account Account = item.(Account)
		if account.Uid == uid {
			uidHasAccess = true
			break
		}
	}

	json.NewEncoder(w).Encode(uidHasAccess)
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
	r.HandleFunc("/api/updateStatus", updateStatus).Methods("POST")
	r.HandleFunc("/api/hasAccess", hasAccess).Methods("POST")
	r.HandleFunc("/api/logAccess", logToDB).Methods("POST")
	r.HandleFunc("/api/open", open).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
