package main

import (
	"code"
	"debugutils"
	"fileutils"
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("./html/dashboard.html"))  //Normalerweise: User sessions (!), Validation server side (returnen einer Seite, die mit Javascript dasselbe form mit error messages zurückgibt (via templates)).Client side: AJAX während dem tippen
var emails map[string]struct{}

func main() {
	initUsers()

	http.HandleFunc("/", handler)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("html"))))
	http.HandleFunc("/usernameExists", userNameExistsHandler)
	http.HandleFunc("/newUser", newUserHandler)

	debugutils.ShutdownRunningServer()
	debugutils.EnableShutdownUtility()
	http.ListenAndServe(":8080", nil)
}

func initUsers() {
	emails = make(map[string]struct{})
	emails["a@a.de"] = struct{}{}
	emails["b@b.de"] = struct{}{}
}

func handler(w http.ResponseWriter, r *http.Request) {
	page, err := fileutils.LoadPage("./html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(page.Body[:]))
}

func userNameExistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		if emailAlreadyUsed(email) {
			fmt.Println("contains")
			fmt.Fprintln(w, "exists")
		}
	}
}

func newUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		if emailAlreadyUsed(email) {
			fmt.Fprintln(w, "exists")
			return
		}
		password := r.FormValue("password")
		emails[email] = struct{}{}

		u := code.NewUser(email, password)
		fmt.Println(u)
		err := templates.ExecuteTemplate(w, "dashboard.html", u)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func emailAlreadyUsed(email string) bool {
	_, contains := emails[email]
	return contains
}
