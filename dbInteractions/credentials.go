package dbInteractions

import "fmt"

const (
	DbHost     = "35.242.197.244"
	DbUser     = "api"
	DbPassword = "TmllbHNTMTZub3Y="
	DbName     = "door_db"
)

var dataSourceString string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", DbUser, DbPassword, DbHost, DbName)
