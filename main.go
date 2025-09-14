package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Lelo88/go-mysql-example/database"
	"github.com/Lelo88/go-mysql-example/handlers"
	"github.com/Lelo88/go-mysql-example/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Connect(func(driverName, dataSourceName string) (*sql.DB, error) {
		return sql.Open(driverName, dataSourceName)
	}, ".env")
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	defer db.Close()

	for {
		fmt.Println("\n MENU")
		fmt.Println("1. List Contacts")
		fmt.Println("2. Get Contact by ID")
		fmt.Println("3. Create Contact")
		fmt.Println("4. Update Contact")
		fmt.Println("5. Delete Contact")
		fmt.Println("6. Exit")
		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			handlers.ListContacts(db)
		case 2:
			var id int
			fmt.Print("Enter Contact ID: ")
			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println("Invalid input, please enter a number.")
				continue
			}
			handlers.GetContactByID(db, id)
		case 3:
			contact := inputContactDetails(3)
			handlers.CreateContact(db, contact)
		case 4:
			contact := inputContactDetails(4)
			handlers.UpdateContact(db, contact)
		case 5:
			var id int
			fmt.Print("Enter Contact ID to delete: ")
			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println("Invalid input, please enter a number.")
				continue
			}
			handlers.DeleteContact(db, id)
		case 6:
			fmt.Println("Exiting...")
			return
		}
	}
}

func inputContactDetails(option int) models.Contact {
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 3 {
		fmt.Print("Enter ID: ")
		var id int
		fmt.Scan(&id)
		contact.ID = id
	}

	if option == 4 {
		fmt.Print("Enter ID of the contact to update: ")
		var id int
		fmt.Scan(&id)
		contact.ID = id
	}

	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Enter Email (or leave blank): ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Enter Phone: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact

}
