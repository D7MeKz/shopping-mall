package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestSetupUserHandler tests the SetupUserHandler function
func TestSetupUserHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/setup-user", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SetupUserHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected HTTP status code 200")
}

// TestStore tests the store function
func TestStore(t *testing.T) {
	// Create a temporary JSON file for testing
	tempFile, err := ioutil.TempFile("", "names.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write test JSON data to the temporary file
	jsonData := `{
		"names": [
			{"username": "Emma"},
			{"username": "Liam"},
			{"username": "Olivia"},
			{"username": "Noah"},
			{"username": "Ava"}
		]
	}`
	if _, err := tempFile.Write([]byte(jsonData)); err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	// Set the path to the temporary file
	originalFilePath := "./list/names.json"
	os.MkdirAll("./list", os.ModePerm)
	defer os.Remove(originalFilePath)
	os.Rename(tempFile.Name(), originalFilePath)

	// Mock database
	mockDB, err := sql.Open("mysql", "user:userpw@tcp(127.0.0.1:3306)/User")
	if err != nil {
		t.Fatalf("Error opening mock database: %v", err)
	}
	defer mockDB.Close()

	err = store(3) // Change count to test partial reading
	assert.NoError(t, err, "Expected no error from store function")
}

// TestParse tests the parse function
func TestParse(t *testing.T) {
	// Create a temporary JSON file for testing
	tempFile, err := ioutil.TempFile("", "names.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write test JSON data to the temporary file
	jsonData := `{
		"names": [
			{"username": "Emma"},
			{"username": "Liam"},
			{"username": "Olivia"},
			{"username": "Noah"},
			{"username": "Ava"}
		]
	}`
	if _, err := tempFile.Write([]byte(jsonData)); err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	// Set the path to the temporary file
	originalFilePath := "./list/names.json"
	os.MkdirAll("./list", os.ModePerm)
	defer os.Remove(originalFilePath)
	os.Rename(tempFile.Name(), originalFilePath)

	users, err := parse(3) // Change count to test partial reading
	assert.NoError(t, err, "Expected no error from parse function")
	assert.Equal(t, 3, len(users), "Expected 3 users to be parsed")
	assert.Equal(t, "Emma", users[0].Username, "Expected first username to be 'Emma'")
	assert.Equal(t, "Liam", users[1].Username, "Expected second username to be 'Liam'")
	assert.Equal(t, "Olivia", users[2].Username, "Expected third username to be 'Olivia'")
}
