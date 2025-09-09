package main

import (
	"fmt"
	"log"

	"github.com/Lelo88/go-mysql-example/database"
	"github.com/Lelo88/go-mysql-example/handlers"
	"github.com/Lelo88/go-mysql-example/models"
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

	fmt.Println("-----")
	
	/* newContact := models.Contact{
		Name:  "John Doe",
		Email: "",
		Phone: "123-456-7890",
	}
	handlers.CreateContact(db, newContact) */

	fmt.Println("-----")
	handlers.ListContacts(db)

	fmt.Println("-----")
	fmt.Println("Update contact with ID 1:")
	updatedContact := models.Contact{
		ID:    1,
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
		Phone: "098-765-4321",
	}
	handlers.UpdateContact(db, updatedContact)
	handlers.ListContacts(db)
}