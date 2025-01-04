package main

import (
    "forum/database"
    "log"
)

func main() {
    database.InitDatabase("./forum.db")

    log.Println("Forum application started successfully!")
}
