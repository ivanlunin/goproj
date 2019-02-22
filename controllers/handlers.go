package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ivanlunin/goproj/models"
)

var (
	// if error occured just send empty json
	emptyJSON = "{}"

	db *models.Database
)

// SetDatabase ...
func SetDatabase(d *models.Database) {
	db = d
}

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO smth with that
	fmt.Fprintf(w, "Home page")
}

// AddPostHandler ...
func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	db.AddPost(title, content)
}

// GetSinglePostHandler ...
func GetSinglePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["id"])
	if err == nil {
		if res, err := db.GetPost(i); err == nil {
			fmt.Fprintf(w, string(res))
			return
		}
	}

	fmt.Fprintf(w, emptyJSON)
}

// GetPostsHandler ...
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	res, err := db.GetAll()
	if err != nil {
		fmt.Fprintf(w, emptyJSON)
	} else {
		fmt.Fprintf(w, string(res))
	}
}
