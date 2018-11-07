package structs

import "time"

// Status struct
type Status struct {
	ID          int       `json:"id"`
	Status      int       `json:"status"`
	Creator     string    `json:"creator"`
	PublishDate time.Time `json:"publishDate"`
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
	Id         int
	Account    string
	AccessDate time.Time
}
