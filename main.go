package main

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
	"fmt"
	// "strconv"
	"net/http"
	"log"
	"html/template"
)

type User struct {
	Uid   int    `json:"uid"`
	Username string `json:"username"`
	Department string `json:"department"`
	Created string `json:"created"`
}

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])
	}
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/about.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])
	}
}

func signin(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])
	}
}

func index(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])
	}
}

func view(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		db, err := sql.Open("mysql", "root:@/test_golang?charset=utf8")
		// query
		results, err := db.Query("select * from userinfo")
		checkErr(err)
		var users []User
		for results.Next() {
			var item_user User
			// for each row, scan the result into our tag composite object
			err = results.Scan(&item_user.Uid, &item_user.Username, &item_user.Department, &item_user.Created)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			users = append(users, User{Uid: item_user.Uid, Username: item_user.Username, Department: item_user.Department,  Created: item_user.Created })
		}
		defer db.Close()

		t, _ := template.ParseFiles("../templates_manajemen/view.html")
		fmt.Println(users)
		data := struct {
			Users []User
		}{
			Users: users,
		}
		t.Execute(w, data)
	} else {
		fmt.Println("oke")
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)              // set router
	http.HandleFunc("/index", index)              // set router
    // http.HandleFunc("/html", html)           // set router
	http.HandleFunc("/login", login)
	http.HandleFunc("/about", about)           // set router
    err := http.ListenAndServe(":9091", nil) // set listen port
    if err != nil {
        log.Fatal("Error running service: ", err)
	}
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}