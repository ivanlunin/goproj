package main

import (
	"fmt"
	"database/sql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
	_ "github.com/lib/pq"
	"os"
)

const (
        DB_USER     = "postgres"
        DB_PASSWORD = "korova228"
        DB_NAME     = "postgres"
)

var db *sql.DB

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Home page")
	t, _ := template.ParseFiles("templates/page.html")
	t.Execute(w, nil)
}

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
        http.Error(w, http.StatusText(405), 405)
        return
    }

	title := r.FormValue("title")
    info := r.FormValue("info")

	fmt.Printf("%3v | %8v\n", title, info)
	
    if title == "" || info == "" {
        http.Error(w, http.StatusText(400), 400)
        return
    }
	
	_, err := db.Exec("INSERT INTO posts(title, info) VALUES($1, $2)", title, info)
	
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
    }
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM users;")
	defer rows.Close()
	checkErr(err)
	
	for rows.Next() {
            var username string
            var password string
			
            _ = rows.Scan(&username, &password)
            
            fmt.Fprintf(w, "%3v | %8v\n", username, password)
    }
	// t, _ := template.ParseFiles("templates/page.html")
    // t.Execute(w, nil)
}

func main() {

	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err = sql.Open("postgres", dbinfo)
	checkErr(err)
    defer db.Close()
	
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/get_users", UsersHandler).Methods("GET")
	router.HandleFunc("/add_post", AddPostHandler).Methods("POST")
	
	fmt.Println("Server started...")
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	http.ListenAndServe(":1337", loggedRouter)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
