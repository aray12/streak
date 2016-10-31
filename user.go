package main

import (
	"fmt"
	db "gopkg.in/dancannon/gorethink.v2"
	"net/http"
)

// CreateUser generates a record in the database
// so user data can be persistent independent of
// of client device
func CreateUser(w http.ResponseWriter, r *http.Request) {
	session, err := db.Connect(db.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		fmt.Println("Connection error ===> ", err.Error())
	} else {
		fmt.Println("Session ===> ", session)
	}

	err = db.DB("hotstreak").Table("users").Insert(map[string]string{
		"id": "john",
		"password": "p455w0rd",
	}).Exec(session)

	fmt.Println("Query error ===> ", err.Error())
}

// Auth authenticates an existing user via a
// username/password combination
func Auth(w http.ResponseWriter, r *http.Request) {

}
