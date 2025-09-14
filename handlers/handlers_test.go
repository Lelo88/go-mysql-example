package handlers

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
}
