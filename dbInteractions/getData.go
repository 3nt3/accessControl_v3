package dbInteractions

import (
	"accessControl_v3/structs"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

// Get data from specified table
func GetData(table string) []interface{} {

	db, err := sql.Open("mysql", dataSourceString)

	// Get the data!!
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s;", table))
	checkErr(err)

	switch table {
	case "accounts":
		var data []interface{}
		for rows.Next() {
			var id string
			var uid string
			var name string
			var tag_name string
			var permission int

			err = rows.Scan(&id, &uid, &name, &tag_name, &permission)
			checkErr(err)

			data = append(data, structs.Account{id, uid, name, tag_name, strconv.Itoa(permission)})
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

			//fmt.Println(publishDate)
			publishDateParsed, _ := time.Parse("2006-01-02 15:04:05", publishDate)

			//fmt.Println(publishDateParsed)

			data = append(data, structs.Status{id, status, creator, publishDateParsed})

		}
		return data
		break
	case "accessLog":
		var data []interface{}

		for rows.Next() {
			var id int
			var uid string
			var userName string
			var publishDate string

			err = rows.Scan(&id, &uid, &userName, &publishDate)
			checkErr(err)

			date, _ := time.Parse("2006-01-02 15:04:05", publishDate)

			data = append(data, structs.Access{id, userName, date})
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
