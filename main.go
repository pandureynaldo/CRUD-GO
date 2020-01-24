package main

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
	"fmt"
	"strconv"
	"net/http"
	"log"
	"html/template"
	"github.com/gorilla/sessions"
	"github.com/gorilla/mux"
	// "reflect"
	"time"
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

type Article struct {
	Id   int    `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Tag string `json:"tag"`
	Created_at string `json:"created_at"`
	Created_by string `json:"created_by"`
	Updated_at string `json:"updated_at"`
	Updated_by string `json:"updated_by"`
	Status string `json:"status"`
}

type Message struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Text string `json:"message"`
	Created_at string `json:"created_at"`
}


func CheckSession(s *sessions.Session) map[string]string{
	// var username, role, errs string

	resp := make(map[string]string)
	if s.Values["username"] != nil {
		resp["username"] = s.Values["username"].(string)
	}

	if s.Values["role"] != nil {
		resp["role"] = s.Values["role"].(string)
	}

	if s.Values["username"] == nil || s.Values["role"] == nil {
		resp["errs"] = "Mohon login"
	}

	return resp

}


func index(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		session, err := store.Get(r, "session-name")
		if err != nil{
			fmt.Println(err)
		}

		header := GetHeader(session)

		// if len(session.Values) == 0 {
		// 	http.Redirect(w, r, "/login", 301)
		// 	return
		// }

		s := CheckSession(session)	
		articles := GetHomeArtikel()
		data := struct {
			Username string
			Role string
			Error string
			Articles []Article
			Header string
		}{
			Username: s["username"],
			Role: s["role"],
			Error: s["errs"],
			Articles: articles,
			Header : header,
		}
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, data)
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
		session.Values["username"] = user.Username
		session.Values["role"] = user.Role
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/about", 301)
		fmt.Println(session)
		fmt.Println(user)
		
	}
}

func logout(w http.ResponseWriter, r *http.Request){
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = nil
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
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
		session.Values["username"] = user.Username
		session.Values["role"] = user.Role
		store.Options = &sessions.Options{
			MaxAge:   60 * 15,
			HttpOnly: true,
		}
	
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index", 301)
		fmt.Println(session)
		fmt.Println(user)
		
	}
}

func GetHomeArtikel() []Article {
	db, err := sql.Open("mysql", "root:@/sa_db?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	results, err := db.Query("select * from article where status = 1")
	checkErr(err)
	var articles []Article
	for results.Next() {
		var item_article Article
		// for each row, scan the result into our tag composite object
		err = results.Scan(&item_article.Id, &item_article.Title, &item_article.Content, &item_article.Created_at, &item_article.Created_by, &item_article.Updated_at, &item_article.Updated_by, &item_article.Status,  &item_article.Tag)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		articles = append(articles, Article{Id: item_article.Id,Title: item_article.Title, Content: item_article.Content, Created_at: item_article.Created_at, Created_by: item_article.Created_by, Updated_at: item_article.Updated_at, Status: item_article.Status, Tag: item_article.Tag})
	}
	defer db.Close()
	return articles;
}

func GetUserArtikel(created_by string) []Article {
	db, err := sql.Open("mysql", "root:@/sa_db?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	results, err := db.Query("select * from article")
	checkErr(err)
	if created_by == "user"{
		results, err = db.Query("select * from article where created_by='"+created_by+"'")
	}
	checkErr(err)
	var articles []Article
	for results.Next() {
		var item_article Article
		// for each row, scan the result into our tag composite object
		err = results.Scan(&item_article.Id, &item_article.Title, &item_article.Content, &item_article.Created_at, &item_article.Created_by, &item_article.Updated_at, &item_article.Updated_by, &item_article.Status,  &item_article.Tag)
		if err != nil {
			// panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(err)
		}
		articles = append(articles, Article{Id: item_article.Id,Title: item_article.Title, Content: item_article.Content, Created_at: item_article.Created_at, Created_by: item_article.Created_by, Updated_at: item_article.Updated_at, Status: item_article.Status, Tag: item_article.Tag})
	}
	defer db.Close()
	return articles;
}

func GetMessagesAdmin() []Message {
	db, err := sql.Open("mysql", "root:@/sa_db?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	results, err := db.Query("select * from contact")
	checkErr(err)
	var messages []Message
	for results.Next() {
		var item_article Message
		// for each row, scan the result into our tag composite object
		err = results.Scan(&item_article.Id, &item_article.Name, &item_article.Email, &item_article.Text, &item_article.Created_at)
		if err != nil {
			// panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(err)
		}
		messages = append(messages, Message{Id: item_article.Id,Name: item_article.Name, Email: item_article.Email, Text: item_article.Text, Created_at: item_article.Created_at})
	}
	defer db.Close()
	return messages;
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


func QueryArticle(id string) Article {
	db, err := sql.Open("mysql", "root:@/sa_db?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	var art = Article{}
	err = db.QueryRow(`
		SELECT id, 
		title, 
		content, 
		tag  
		FROM article WHERE id=?
		`, id).
		Scan(
			&art.Id,
			&art.Title,
			&art.Content,	
			&art.Tag,
		)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(art)
	defer db.Close()
	return art
}

func GetHeader(s *sessions.Session) string{
	resp := make(map[string]string)
	if s.Values["username"] != nil {
		resp["username"] = s.Values["username"].(string)
	}
	var template string
	// if s.Values["role"] != nil {
	// 	resp["role"] = s.Values["role"].(string)
	// }

	// if s.Values["username"] == nil || s.Values["role"] == nil {
	// 	resp["errs"] = "Mohon login"
	// }
	if s.Values["username"] != nil && s.Values["role"] != nil{
		template = `<li class="active"><a href="#">Home</a></li>
		<li><a href="/about">Tentang Kami</a></li>
		<li><a href="/contact">Kontak Kami</a></li>
		<li><a href="/login">Login</a></li>`
	}else{
		template = `<li class="active"><a href="#">Home</a></li>
		<li><a href="/about">Tentang Kami</a></li>
		<li><a href="/contact">Kontak Kami</a></li>
		<li><a href="/login">Login</a></li>`
	}
	return template
}

func harticle(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	session, err := store.Get(r, "session-name")
	if err != nil{
		fmt.Println(err)
	}

	// if len(session.Values) == 0 {
	// 	http.Redirect(w, r, "/login", 301)
	// 	return
	// }

	s := CheckSession(session)	
	if s["errs"] != ""{
		http.Redirect(w, r, "/login", 301)
	}
	articles := GetUserArtikel(s["username"])
	// for index, element := range articles {
	// 	if 
	// }
	if r.Method == "GET" {
		data := struct {
			Username string
			Role string
			Error string
			Articles []Article
		}{
			Username: s["username"],
			Role: s["role"],
			Error: s["errs"],
			Articles : articles,
			
		}
		t, err := template.ParseFiles("templates/article.html")
		if err != nil{
			fmt.Println(err)
		}
		t.Execute(w, data)
	} else {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])
	}
}


func about(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	session, err := store.Get(r, "session-name")
	if err != nil{
		fmt.Println(err)
	}

	// if len(session.Values) == 0 {
	// 	http.Redirect(w, r, "/login", 301)
	// 	return
	// }

	s := CheckSession(session)	

	if r.Method == "GET" {
		data := struct {
			Username string
			Role string
			Error string
		}{
			Username: s["username"],
			Role: s["role"],
			Error: s["errs"],
		}
		t, err := template.ParseFiles("templates/about.html")
		if err != nil{
			fmt.Println(err)
		}
		t.Execute(w, data)
	} else {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])
	}
}

func contact(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	session, err := store.Get(r, "session-name")
	if err != nil{
		fmt.Println(err)
	}

	// if len(session.Values) == 0 {
	// 	http.Redirect(w, r, "/login", 301)
	// 	return
	// }

	s := CheckSession(session)	
	fmt.Println(s)
	var messages []Message
	
	if r.Method == "GET" {
		var files_name string
		if s["role"] == "admin"{
			messages = GetMessagesAdmin()
			files_name = "contact_admin.html"
		}else{
			files_name = "contact.html"
		}
		t, _ := template.ParseFiles("templates/"+files_name)
		data := struct {
			Username string
			Role string
			Error string
			Messages []Message
		}{
			Username: s["username"],
			Role: s["role"],
			Error: s["errs"],
			Messages : messages,
		}
		t.Execute(w, data)
	} else {
		r.ParseForm()
		session, err := store.Get(r, "session-name")
		if err != nil{
			fmt.Println(err)
			return
		}
		s := CheckSession(session)
		if s["errs"] != "" {
			fmt.Println(s["errs"])
			http.Redirect(w, r, "/login", 301)
			return
		}
		db, err := sql.Open("mysql", "root:@/sa_db?charset=utf8")
		checkErr(err)

		// insert
		stmt, err := db.Prepare("insert into `contact` values (null, ?, ?, ?, ?)")
		checkErr(err)

		currentTime := time.Now()

		res, err := stmt.Exec(r.Form["name"][0], r.Form["email"][0], r.Form["message"][0], currentTime.Format("2006-01-02 15:04:05"))
		if err != nil{
			fmt.Println(err)
		}
		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
		http.Redirect(w, r, "/contact", 301)
	}
}

func status(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	session, err := store.Get(r, "session-name")
	if err != nil{
		fmt.Println(err)
		return
	}
	s := CheckSession(session)
	if s["errs"] != "" {
		fmt.Println(s["errs"])
		http.Redirect(w, r, "/login", 301)
		return
	}
	if r.Method == "GET" {
		// keys := r.URL.Query()
		vars := mux.Vars(r)
		id := vars["id"]
		value := vars["value"]
		// status := keys.Get("status")
		db, err := sql.Open("mysql", "root:@/sa_db?charset=utf8")
		// update
		stmt, err := db.Prepare("update `article` set status=? where id=?")
		checkErr(err)
		// // id, err := res.LastInsertId()
		// // checkErr(err)
		
		if err != nil{
			fmt.Println(err)
		}
		res, err := stmt.Exec(value,id)
		checkErr(err)

		affect, err := res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)

		defer db.Close()
		
	} 
	http.Redirect(w, r, "/article", 301)
}


func edit(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	session, err := store.Get(r, "session-name")
	if err != nil{
		fmt.Println(err)
	}

	// if len(session.Values) == 0 {
	// 	http.Redirect(w, r, "/login", 301)
	// 	return
	// }

	s := CheckSession(session)
	if r.Method == "GET" {
		keys := r.URL.Query()
		id := keys.Get("id")
		articles := QueryArticle(id)
		fmt.Println(articles)
		data := struct {
			Username string
			Role string
			Error string
			Id string
			Title string
			Content string
			Tag string
		}{
			Username: s["username"],
			Role: s["role"],
			Error: s["errs"],
			Id : strconv.Itoa(articles.Id),
			Title : articles.Title,
			Content : articles.Content,
			Tag : articles.Tag,
			
		}
		t, err := template.ParseFiles("templates/edit.html")
		if err != nil{
			fmt.Println(err)
		}
		t.Execute(w, data)
	} else {
		r.ParseForm()
		db, err := sql.Open("mysql", "root:@/sa_db?charset=utf8")
		// update
		stmt, err := db.Prepare("update `article` set title=?, content=?, tag=? where id=?")
		checkErr(err)
		// // id, err := res.LastInsertId()
		// // checkErr(err)
		
		article_id, errs := strconv.Atoi(r.Form["Aid"][0])
		
		if errs != nil{
			fmt.Println(errs)
		}

		fmt.Println(article_id)
		res, err := stmt.Exec(r.Form["title"][0], r.Form["content"][0], r.Form["tag"][0], article_id)
		checkErr(err)

		affect, err := res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)

		defer db.Close()
		http.Redirect(w, r, "/article", 301)
	}
}

func addarticle(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method: ", r.Method) // get request method
	session, err := store.Get(r, "session-name")
	if err != nil{
		fmt.Println(err)
		http.Redirect(w, r, "/article", 301)
	}



	// if len(session.Values) == 0 {
	// 	http.Redirect(w, r, "/login", 301)
	// 	return
	// }

	s := CheckSession(session)	
	if s["errs"] != ""{
		http.Redirect(w, r, "/login", 301)
	}
	articles := GetUserArtikel(s["username"])
	if r.Method == "GET" {
		data := struct {
			Username string
			Role string
			Error string
			Articles []Article
		}{
			Username: s["username"],
			Role: s["role"],
			Error: s["errs"],
			Articles : articles,
			
		}
		t, err := template.ParseFiles("templates/add.html")
		if err != nil{
			fmt.Println(err)
		}
		t.Execute(w, data)
	} else {
		r.ParseForm()
		db, err := sql.Open("mysql", "root:@/sa_db?charset=utf8")
		checkErr(err)

		// insert
		stmt, err := db.Prepare("insert into `article` values (null, ?, ?, ?, ?, ?, ?, ?, ?)")
		checkErr(err)

		currentTime := time.Now()

		res, err := stmt.Exec(r.Form["title"][0], r.Form["content"][0], currentTime.Format("2006-01-02 15:04:05"),s["username"], s["username"], currentTime.Format("2006-01-02 15:04:05"), 0, r.Form["tag"][0])
		if err != nil{
			fmt.Println(err)
		}
		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
		http.Redirect(w, r, "/article", 301)
	}
}

func home(w http.ResponseWriter, r *http.Request){
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
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/login", login)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/index", index)
	router.HandleFunc("/about", about)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/article", harticle)
	router.HandleFunc("/add", addarticle)
	router.HandleFunc("/edit", edit)
	router.HandleFunc("/status/{id}/{value:[0-1]+}", status)
	// http.Handle("/static/", http.StripPrefix("./static/*", http.FileServer(http.Dir("./static"))))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// router.PathPrefix("/static").Handler(http.FileServer(http.Dir("static")))
	// http.HandleFunc("/", index)              // set router
	// http.HandleFunc("/index", index)              // set router
    // http.HandleFunc("/html", html)           // set router
	// http.HandleFunc("/login", login)
	// http.HandleFunc("/logout", logout)
	// http.HandleFunc("/about", about)  
	// http.HandleFunc("/contact", contact)           // set router
	
    err := http.ListenAndServe(":9091", router) // set listen port
    if err != nil {
        log.Fatal("Error running service: ", err)
	}
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}