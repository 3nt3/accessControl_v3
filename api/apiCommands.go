package api

import (
	"accessControl_v3/dbInteractions"
	"accessControl_v3/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Get all statuses
func GetStatuses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbInteractions.GetData("statusLog"))
}

// Get current status
func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := dbInteractions.GetData("statusLog")

	json.NewEncoder(w).Encode(data[len(data)-1])

}

// Update current Status
func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newStatus structs.Status
	_ = json.NewDecoder(r.Body).Decode(&newStatus)

	newStatus.PublishDate = time.Now()
	fmt.Println(newStatus)

	data := []interface{}{newStatus.Status, newStatus.Creator, newStatus.PublishDate}
	dbInteractions.InsertData("statusLog", data)
	json.NewEncoder(w).Encode(dbInteractions.GetData("statusLog")[len(dbInteractions.GetData("statusLog"))-1])
}

// Check if UID has access
func HasAccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonData map[string]string
	_ = json.NewDecoder(r.Body).Decode(&jsonData)

	uid := jsonData["uid"]
	var uidHasAccess bool

	for _, item := range dbInteractions.GetData("accounts") {
		var account structs.Account = item.(structs.Account)

		if account.Uid == uid {
			if account.Permission == "99" || account.Permission == "100" {
				uidHasAccess = true
				break
			} else {
				accesses := 0
				for _, item := range dbInteractions.GetData("accessLog") {
					var access structs.Access = item.(structs.Access)
					var isToday bool

					accessYear, accessDay := access.AccessDate.Year(), access.AccessDate.YearDay()
					now := time.Now()

					if accessYear == now.Year() && accessDay == now.YearDay() {
						isToday = true
					}

					if access.Account == account.Name && isToday {
						accesses++
					}
				}
				permission, _ := strconv.Atoi(account.Permission)

				if accesses < permission {
					uidHasAccess = true
					fmt.Printf("%s - %s hasPermission\nPossible accesses today: %d\n", account.Uid, account.Name, permission-accesses)
				} else {
					fmt.Printf("%s - %s !hasPermission\nPossible accesses today: %d\n", account.Uid, account.Name, permission-accesses)
				}
			}
		}
	}

	json.NewEncoder(w).Encode(uidHasAccess)
}

// Open the door
func Open(w http.ResponseWriter, r *http.Request) {
	// I actually do not have any clue how to open the door
}

// Log to DB
func LogAccess(w http.ResponseWriter, r *http.Request) {

}

func TestConn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("it works!")
}
