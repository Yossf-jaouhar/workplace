package backend

import "database/sql"

var DB *sql.DB


type user struct {
	
}

type Err struct {
	Type string
	Code int
}