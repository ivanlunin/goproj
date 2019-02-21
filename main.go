package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ivanlunin/goproj/controllers"
	"github.com/ivanlunin/goproj/models"
)

func main() {
	db := &models.Database{}
	db.Init()

	controllers.SetDatabase(db)

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/api/v1/get_posts", controllers.GetPostsHandler).Methods("GET")
	router.HandleFunc("/api/v1/get_post/{id:[0-9]+}", controllers.GetSinglePostHandler).Methods("GET")
	router.HandleFunc("/api/v1/add_post", controllers.AddPostHandler).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	log.Fatal(srv.ListenAndServe())
}
