package controllers

import (
	"fmt"
	postsRepository "go-web/repositories"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := postsRepository.GetPosts()

	fmt.Fprintf(w, string(posts))
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	postsRepository.CreatePost()
}
