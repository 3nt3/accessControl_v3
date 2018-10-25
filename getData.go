package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go/types"
	"strconv"
	"time"
)

const DB_HOST = "host"
const DB_USR = "account"
const DB_PASSWORD = "password"
const DB_NAME = "name"

var dataSorceString string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", DB_USR, DB_PASSWORD, DB_HOST, DB_NAME)

// Get accounts
func GetData(table string) types.Slice {

	db, err := sql.Open("mysql", dataSorceString)
	if err != nil {
		fmt.Println("Connection Failed:", err)
	} else {
		fmt.Println("Connected!")
	}

	table = "statusLog"

	// Get the data!!
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s;", table))
	checkErr(err)

	fmt.Println(table)

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
		return data
	case "statusLog":
		data := []Status{}

		for rows.Next() {
			var id int
			var status int
			var creator string
			var publishDate string
			err = rows.Scan(&id, &status, &creator, &publishDate)
			checkErr(err)

			fmt.Println(data)

			data = append(data, Status{strconv.Itoa(id), strconv.Itoa(status), creator, publishDate})
			fmt.Printf("After: %s\n", data)

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

			date, _ := time.Parse("2006-01-02 15:04:05", publishDate)

			data = append(data, Access{id, account, date})
			fmt.Println(data)

		}
	default:
		fmt.Println("The specified table does not exist!")
	}

	//fmt.Println(Accounts)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
