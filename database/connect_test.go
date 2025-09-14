package database

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func setEnvVars() {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "password")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "db_contacts")
}

func Test_Connect(t *testing.T) {
	t.Run("Successful Connection", func(t *testing.T) {
		// Set environment variables
		setEnvVars()

		mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer mockDB.Close()

		mock.ExpectPing().WillReturnError(nil)

		db, err := Connect(func(driverName, dataSourceName string) (*sql.DB, error) {
			return mockDB, nil
		})
		require.NoError(t, err, "unexpected error")
		require.NotNil(t, db, "db should not be nil")
	})

	t.Run("Error Loading .env", func(t *testing.T) {
		// Simulate missing environment variables
		os.Clearenv()

		mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer mockDB.Close()

		mock.ExpectPing().WillReturnError(nil)

		db, err := Connect(func(driverName, dataSourceName string) (*sql.DB, error) {
			return mockDB, nil
		}, "nonexistent.env")
		require.Error(t, err, "expected an error but got none")
		require.Nil(t, db, "db should be nil")
	})

	t.Run("Error Connecting to Database", func(t *testing.T) {
		// Set environment variables
		setEnvVars()

		mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer mockDB.Close()

		mock.ExpectPing().WillReturnError(errors.New("connection error"))

		db, err := Connect(func(driverName, dataSourceName string) (*sql.DB, error) {
			return mockDB, nil
		})
		require.Error(t, err, "expected an error but got none")
		require.Nil(t, db, "db should be nil")
	})

	t.Run("Error Loading Invalid .env", func(t *testing.T) {
		// Pass an invalid .env path
		os.Clearenv()

		mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer mockDB.Close()

		mock.ExpectPing().WillReturnError(nil)

		db, err := Connect(func(driverName, dataSourceName string) (*sql.DB, error) {
			return mockDB, nil
		}, "invalid.env")
		require.Error(t, err, "expected an error but got none")
		require.Nil(t, db, "db should be nil")
	})

	t.Run("Missing Environment Variables", func(t *testing.T) {
		// Clear environment variables
		os.Clearenv()

		mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		require.NoError(t, err, "An error was not expected when opening a stub database connection")
		defer mockDB.Close()

		mock.ExpectPing().WillReturnError(nil)

		db, err := Connect(func(driverName, dataSourceName string) (*sql.DB, error) {
			return mockDB, nil
		})
		require.Error(t, err, "expected an error but got none")
		require.Nil(t, db, "db should be nil")
	})

	t.Run("Nil DB Constructor", func(t *testing.T) {
		// Set environment variables
		setEnvVars()

		// Call Connect with nil dbConstructor
		db, err := Connect(nil)
		require.Error(t, err, "expected an error but got none")
		require.Nil(t, db, "db should be nil")
	})
}
