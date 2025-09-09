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

func GetContactByID(db *sql.DB, id int) {
	// Implementation for getting a contact by ID
	query := "SELECT * FROM contact WHERE id = ?"
	
	row := db.QueryRow(query, id)
	contact := models.Contact{}

	var valueEmail sql.NullString
	
	err := row.Scan(&contact.ID, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No contact found with the given ID.")
			return
		}
		log.Fatal("Could not scan row:", err)
	}
	
	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "No Email"
	}

	fmt.Println("Contact Details:")
	fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
}

func CreateContact(db *sql.DB, contact models.Contact) error {
	// Implementation for creating a new contact
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"
	
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		return fmt.Errorf("could not execute insert query: %v", err)
	}

	fmt.Println("Contact created successfully!")
	return nil
}

func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.ID)
	if err != nil {
		log.Fatal("Could not execute update query:", err)
	}

	fmt.Println("Contact updated successfully!")
}

func DeleteContact(db *sql.DB, id int) {
	query := "DELETE FROM contact WHERE id = ?"

	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("Could not execute delete query:", err)
	}

	fmt.Println("Contact deleted successfully!")
}