package main

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
	"fmt"
	// "strconv"
	"net/http"
	"log"
	"html/template"
	"github.com/gorilla/sessions"
	// "reflect"
)

var (
    // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
    key = []byte("this-is-secret")
    store = sessions.NewCookieStore(key)
)

type User struct {
	Id   int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Role string `json:"role"`
	Status string `json:"status"`
}

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		username := r.Form["username"][0]
		fmt.Println("Password: ", r.Form["password"])
		password := r.Form["password"][0]
		user := QueryUser(username,password);
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
	}
}


func QueryUser(username string, password string) User {
	db, err := sql.Open("mysql", "root:@/sa_db?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	var users = User{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		password, 
		role  
		FROM users WHERE username=? and password=? and status=?
		`, username, password, 1).
		Scan(
			&users.Id,
			&users.Username,
			&users.Password,
			&users.Role,
		)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(users)
	defer db.Close()
	return users
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

func contact(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/contact.html")
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

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)              // set router
	http.HandleFunc("/index", index)              // set router
    // http.HandleFunc("/html", html)           // set router
	http.HandleFunc("/login", login)
	http.HandleFunc("/about", about)  
	http.HandleFunc("/contact", contact)           // set router
	
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