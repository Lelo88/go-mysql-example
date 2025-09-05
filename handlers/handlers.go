package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Lelo88/go-mysql-example/models"
)

func ListContacts(db *sql.DB) {
	// Implementation for listing contacts
	query := "SELECT * FROM contact"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Could not execute query:", err)
	}
	defer rows.Close()

	fmt.Println("Contacts:")

	for rows.Next() {
		contact := models.Contact{}

		var valueEmail sql.NullString

		err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Phone)
		if err != nil {
			log.Fatal("Could not scan row:", err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "No Email"
		}

		fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	}
}