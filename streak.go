package main

import (
	"fmt"
	"net/http"
	db "gopkg.in/dancannon/gorethink.v2"
)

func handler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}