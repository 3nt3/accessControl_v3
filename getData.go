package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go/types"
	"strconv"
)

const DB_HOST = "35.242.197.244"
const DB_USR = "root"
const DB_PASSWORD = "TmllbHNTMTZub3Y="
const DB_NAME = "door_db"

var dataSorceString string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", DB_USR, DB_PASSWORD, DB_HOST, DB_NAME)

func main() {
	GetData()
}

// Get accounts
func GetData(table string) types.Slice {

	data := types.Slice{}

	db, err := sql.Open("mysql", dataSorceString)
	if err != nil {
		fmt.Println("Connection Failed:", err)
	} else {
		fmt.Println("Connected!")
	}

	// Get the data!!
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s;", table))
	checkErr(err)

	switch table {
	case "accounts":
		data := []Account{}

		for rows.Next() {
			var id int
			var uid string
			var name string
			var tag_name string
			var permission int
			err = rows.Scan(&id, &uid, &name, &tag_name, &permission)
			checkErr(err)

			data = append(data, Account{strconv.Itoa(id), uid, name, tag_name, strconv.Itoa(permission)})
		}
	case "statusLog":
		data := []Status{}

		for rows.Next() {
			var id int
			var status int
			var creator string
			var publishDate string
			err = rows.Scan(&id, &status, &creator, &publishDate)
			checkErr(err)

			data = append(data, Status{strconv.Itoa(id), strconv.Itoa(status), creator, publishDate})
		}
	case "accessLog":
		data := []Access{}

		for rows.Next() {
			var id int
			var status int

			var userID int
			var userUID int
			var userName string
			var userTagName string
			var userPermission string

			var publishDate string
			err = rows.Scan(&id, &status, &userID, &userUID, &userName, &userTagName, &userPermission, &publishDate)
			checkErr(err)

			account := Account{strconv.Itoa(userID), strconv.Itoa(userUID), userName, userTagName, userPermission}

			data = append(data, Access{id, account, publishDate})
		}

	}

	//fmt.Println(Accounts)

	db.Close()

	return data
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
