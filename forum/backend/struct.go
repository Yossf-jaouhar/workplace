package backend

import "database/sql"

var DB *sql.DB


type Err struct {
	Type string
	Code int
}