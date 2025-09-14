package handlers

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Lelo88/go-mysql-example/models"
	"github.com/stretchr/testify/require"
)

func Test_ListContacts(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow(1, "John Doe", "john.doe@example.com", "123456789").
			AddRow(2, "Jane Smith", "jane.smith@example.com", "987654321")

		mock.ExpectQuery("SELECT \\* FROM contact").WillReturnRows(rows)

		err = ListContacts(db)
		require.NoError(t, err, "unexpected error")

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "there were unfulfilled expectations")
	})

	t.Run("Query Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		mock.ExpectQuery("SELECT \\* FROM contact").WillReturnError(errors.New("query error"))

		err = ListContacts(db)
		require.Error(t, err, "expected an error but got none")
		require.EqualError(t, err, "could not execute query: query error", "unexpected error message")
	})

	t.Run("Row Scan Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow(1, nil, "invalid-email", "123456789")

		mock.ExpectQuery("SELECT \\* FROM contact").WillReturnRows(rows)

		err = ListContacts(db)
		require.Error(t, err, "expected an error but got none")
	})

	t.Run("Valid Email", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow(1, "John Doe", "john.doe@example.com", "123456789")

		mock.ExpectQuery("SELECT \\* FROM contact").WillReturnRows(rows)

		err = ListContacts(db)
		require.NoError(t, err, "unexpected error")

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "there were unfulfilled expectations")
	})

	t.Run("Invalid Email Format", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow(1, "John Doe", sql.NullString{String: "invalid-email", Valid: true}, "123456789")

		mock.ExpectQuery("SELECT \\* FROM contact").WillReturnRows(rows)

		err = ListContacts(db)
		require.Error(t, err, "expected an error for invalid email format but got none")
		require.EqualError(t, err, "invalid email format: invalid-email", "unexpected error message")

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "there were unfulfilled expectations")
	})

	t.Run("No Email Assigned", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow(1, "John Doe", sql.NullString{String: "", Valid: false}, "123456789")

		mock.ExpectQuery("SELECT \\* FROM contact").WillReturnRows(rows)

		err = ListContacts(db)
		require.NoError(t, err, "unexpected error")

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "there were unfulfilled expectations")
	})

	t.Run("Email is NULL", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		row := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow(1, "John Doe", sql.NullString{String: "", Valid: false}, "123456789")

		mock.ExpectQuery("SELECT \\* FROM contact").WillReturnRows(row)

		err = ListContacts(db)
		require.NoError(t, err, "unexpected error")

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "there were unfulfilled expectations")
	})
}

func TestGetContactByID(t *testing.T) {
	t.Run("Contact Found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		row := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow(1, "John Doe", "john.doe@example.com", "123456789")

		mock.ExpectQuery("SELECT \\* FROM contact WHERE id = \\?").WithArgs(1).WillReturnRows(row)

		err = GetContactByID(db, 1)
		require.NoError(t, err, "unexpected error")
	})

	t.Run("Contact Not Found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		mock.ExpectQuery("SELECT \\* FROM contact WHERE id = \\?").WithArgs(2).WillReturnError(sql.ErrNoRows)

		err = GetContactByID(db, 2)
		require.Error(t, err, "expected an error but got none")
		require.EqualError(t, err, "no contact found with the given ID: 2", "unexpected error message")
	})

	t.Run("Scan Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		row := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow("invalid", "John Doe", "john.doe@example.com", "123456789")

		mock.ExpectQuery("SELECT \\* FROM contact WHERE id = \\?").WithArgs(1).WillReturnRows(row)

		err = GetContactByID(db, 1)
		require.Error(t, err, "expected an error but got none")
	})
}

func Test_CreateContact(t *testing.T) {
	t.Run("Successful Insertion", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		contact := models.Contact{
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Phone: "123-456-7890",
		}
		
		mock.ExpectExec("INSERT INTO contact \\(name, email, phone\\) VALUES \\(\\?, \\?, \\?\\)").
			WithArgs(contact.Name, contact.Email, contact.Phone).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = CreateContact(db, contact)
		require.NoError(t, err, "unexpected error")

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "there were unfulfilled expectations")
	})

	t.Run("Insertion Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer db.Close()

		contact := models.Contact{
			Name:  "John Doe",
			Email: "john.doe@example.com",		
			Phone: "123-456-7890",
		}

		mock.ExpectExec("INSERT INTO contact \\(name, email, phone\\) VALUES \\(\\?, \\?, \\?\\)").
			WithArgs(contact.Name, contact.Email, contact.Phone).
			WillReturnError(errors.New("insertion error"))

		err = CreateContact(db, contact)
		require.Error(t, err, "expected an error but got none")
		require.EqualError(t, err, "could not execute insert query: insertion error", "unexpected error message")

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "there were unfulfilled expectations")
	})
}