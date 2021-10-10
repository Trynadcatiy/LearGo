package main

import (
	"fmt"
	"net/http"
)

func homePage(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "Go is super easy!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts", contactsPage)
	http.ListenAndServe(":8000", nil)
}

func contactsPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Contacts page!")
}

func main() {
	handleRequest()
}
