package main

import (
	"log"

	"github.com/Lelo88/go-mysql-example/database"
	"github.com/Lelo88/go-mysql-example/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	defer db.Close()

	handlers.ListContacts(db)
}