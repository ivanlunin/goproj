package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ivanlunin/goproj/models"
)

var (
	db *models.Database
)

// SetDatabase ...
func SetDatabase(d *models.Database) {
	db = d
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
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return
		}
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// GetPostsHandler ...
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	res, err := db.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}
