package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func main() {
	// Explicitly declaring err here to avoid using := syntax in the
	// db connection statement. Using := will create a new db variable
	// limited to this scope instead of initializing the global db var.
	var err error

	db, err = pgxpool.Connect(context.Background(), os.Getenv("DB_CONN"))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/login", Login)
	http.HandleFunc("/signup", Signup)

	log.Fatalln(http.ListenAndServe(":8080", nil))

}

type Credentials struct {
	Username string `json:"username", db:"username"`
	Password string `json:"password", db:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	inputInfo := &Credentials{}
	err = json.NewDecoder(r.Body).Decode(inputInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	creds := db.QueryRow(context.Background(), "SELECT password FROM credential WHERE username=$1", inputInfo.Username)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	storedInfo := &Credentials{}
	err = creds.Scan(&storedInfo.Password)
	if err != nil {
		//if username not found
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "User does not exist: %v\n", err)
		} else {
			//other error
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	//Compare the hash of password from user input with the hashed password in database
	err = bcrypt.CompareHashAndPassword([]byte(storedInfo.Password), []byte(inputInfo.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Incorrect password: %v\n", err)
		return
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var err error
	inputInfo := &Credentials{}
	err = json.NewDecoder(r.Body).Decode(inputInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputInfo.Password), bcrypt.DefaultCost)
	_, err = db.Query(context.Background(), "insert into users values ($1, $2)", inputInfo.Username, string(passwordHash))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
