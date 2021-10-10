package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name  string
	Age   uint16
	Money int16
	AvgGrades,
	Happiness float64
	Hobbies []string
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. "+
		"He is %d age and he has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(writer http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Gym", "Coding"}}
	//fmt.Fprintf(writer, `<h1>Main text</h1>
	//<b>Main text</b>`)

	tmpl, _ := template.ParseFiles("templates/homepage.html")
	tmpl.Execute(writer, bob)
}

func contactsPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts", contactsPage)
	http.ListenAndServe(":8000", nil)
}

func main() {
	//var bob User = ....
	//bob := User{name: "Bob", age: 25, money: -50, avgGrades: 4.2, happiness: 0.8}

	handleRequest()
}
