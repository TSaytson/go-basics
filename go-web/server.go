package main

import (
	postsController "go-web/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	r.HandleFunc("/getPosts", postsController.GetPosts).Methods("GET")

	r.HandleFunc("/createPost", postsController.CreatePost).Methods("POST")

	handler := cors.Handler(r)
	http.ListenAndServe(":3001", handler)
}
