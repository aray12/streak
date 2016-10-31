package main

import "net/http"

func main() {
	http.HandleFunc("/", CreateUser)
	http.ListenAndServe(":8081", nil)
}