package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// This is a placeholder for the main function.
	fmt.Println("Hello, World!")

	dns := "root:@tcp(localhost:3306)/db_contacts"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database successfully!")
	defer db.Close()
}