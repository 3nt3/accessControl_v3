package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var dataSorceString string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", DB_USR, DB_PASSWORD, DB_HOST, DB_NAME)

// Get accounts
func GetData(table string) []interface{} {

	db, err := sql.Open("mysql", dataSorceString)

	// Get the data!!
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s;", table))
	checkErr(err)

	switch table {
	case "accounts":
		var data []interface{}
		for rows.Next() {
			var id int
			var uid string
			var name string
			var tag_name string
			var permission int

			err = rows.Scan(&id, &uid, &name, &tag_name, &permission)
			checkErr(err)

			data = append(data, Account{strconv.Itoa(id), uid, name, tag_name, strconv.Itoa(permission)})
			//fmt.Println(data)
		}
		return data
		break
	case "statusLog":
		var data []interface{}

		for rows.Next() {
			var id int
			var status int
			var creator string
			var publishDate string
			err = rows.Scan(&id, &status, &creator, &publishDate)
			checkErr(err)

			data = append(data, Status{id, strconv.Itoa(status), creator, publishDate})

		}
		return data
		break
	case "accessLog":
		var data []interface{}

		for rows.Next() {
			var id int
			var status int

			var userName string

			var publishDate string

			err = rows.Scan(&id, &status, &userName, &publishDate)
			checkErr(err)

			date, _ := time.Parse("2006-01-02 15:04:05", publishDate)

			data = append(data, Access{id, userName, date})
		}
		return data
		break
	default:
		fmt.Println("The specified table does not exist!")
		return nil
		break
	}

	//fmt.Println(Accounts)

	db.Close()
	return nil
}

func InsertData(table string, data []interface{}) bool {

	// Establish connection
	db, err := sql.Open("mysql", dataSorceString)
	if err != nil {
		fmt.Println("Connection Failed:", err)
		return false
	}

	// Actual insert
	switch table {
	case "accounts":
		stmt, err := db.Prepare(fmt.Sprintf("INSERT %s SET uid=?,name=?,tag_name=?,permission=?", table))
		checkErr(err)

		_, err = stmt.Exec(data[0], data[1], data[2], data[3])
		checkErr(err)
		return true

	case "statusLog":
		stmt, err := db.Prepare(fmt.Sprintf("INSERT %s SET status=?,creator=?", table))
		checkErr(err)

		_, err = stmt.Exec(data[0], data[1])
		checkErr(err)
		return true

	case "accessLog":
		stmt, err := db.Prepare(fmt.Sprintf("INSERT %s SET uid=?,name=?,publishDate=?", table))
		checkErr(err)

		_, err = stmt.Exec(data[0], data[1], data[2])
		checkErr(err)
		return true
	}
	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
