package main

import (
	"encoding/json"
	"fmt"
	"net/http" // import package net-http เข้ามา
)

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Warathat",
		Lastname:  "Pant",
		Code:      1997,
		Phone:     "0812345678",
	}
	json.NewEncoder(w).Encode(addBook)
}

func homePage(w http.ResponseWriter, r *http.Request) { // (1)
	fmt.Fprint(w, "Welcome to the HomePage!") // (2)
}
func handleRequest() { // (3)
	http.HandleFunc("/", homePage) // (4)
	http.HandleFunc("/getAddress", getAddressBookAll)
	http.ListenAndServe(":8080", nil) // (5)
}

func main() {
	handleRequest() // ต้องนำ handleRequest มาใส่ใน main ด้วยนะครับ
}
