package dbInteractions

import "fmt"

const (
	DB_HOST     = "35.242.197.244"
	DB_USR      = "root"
	DB_PASSWORD = "s"
	DB_NAME     = "door_db"
)

var dataSorceString string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", DB_USR, DB_PASSWORD, DB_HOST, DB_NAME)
