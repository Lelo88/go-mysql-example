package main

import (
	"fmt"
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
	fmt.Println("-----")
	fmt.Println("Get contact with ID 1:")
	handlers.GetContactByID(db, 1)
}