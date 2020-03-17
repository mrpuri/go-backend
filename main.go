package main

import (
	"fmt"
	"net/http"

	connection "backend/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"encoding/json"
)
type User struct{
	user string
	password string
}

var r = mux.NewRouter()

func main() {

	r.HandleFunc("/user/{username}/{password}", addUser).Methods("POST")
	r.HandleFunc("/getUsers", getUsers).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	{ // Insert a new user

// 		decoder := json.NewDecoder(request.Body)
// err := decoder.Decode(&userDetails) 

		db := connection.DbConn()
		vars := mux.Vars(r)
		username := vars["username"]
		password := vars["password"]

		result, err := db.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, username, password)
		if err != nil {
			fmt.Println(err)
		}

		id, err := result.LastInsertId()

		fmt.Println(id)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db := connection.DbConn()
	selDB, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(selDB.Columns)
}
