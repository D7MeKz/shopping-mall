package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"os"
)

func SetupUserHandler(w http.ResponseWriter, r *http.Request) {
	// METHOD : GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// store data
	err := store(10)
	if err != nil {
		http.Error(w, "Error storing data", http.StatusInternalServerError)
		return
	}

}

func store(cnt int) error {
	dsn := "user:userpw@tcp(127.0.0.1:3306)/User"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return err
	}

	defer func(db *sql.DB) {
		derr := db.Close()
		if derr != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}(db)

	// parse data from names.json
	users, err := parse(cnt)
	if err != nil {
		log.Fatalf("Error parsing data: %v", err)
		return err
	}

	// insert to user
	for _, user := range users {
		derr := insert(db, user.Username)
		if derr != nil {
			return err
		}
	}

	return nil
}

func insert(db *sql.DB, username string) error {
	query := "INSERT INTO Member (username) VALUES (?)"
	_, err := db.Exec(query, username)
	if err != nil {
		log.Printf("Error storing data in database: %v", err)
		return err
	}
	return nil
}

type User struct {
	Username string `json:"username"`
}

// UserList represents the structure of the JSON file
type UserList struct {
	Names []User `json:"names"`
}

// parse() function is parsing names from names.json
func parse(cnt int) ([]User, error) {
	f, err := os.Open("./list/names.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	byteValue, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var data UserList
	if merr := json.Unmarshal(byteValue, &data); merr != nil {
		log.Fatalf("Error unmarshalling data: %v", merr)
	}

	if cnt > len(data.Names) {
		cnt = len(data.Names)
	}

	return data.Names[:cnt], nil
}
