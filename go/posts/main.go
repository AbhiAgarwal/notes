// http://www.gorillatoolkit.org/pkg/mux

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home")
}

func PostIndexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post index")
}

func PostCreateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post create")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post edit")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(rw, "post show "+id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post delete")
}

func main() {
	// Initialization of mux
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", HomeHandler)

	// Collection of Posts
	posts := r.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(PostIndexHandler)
	posts.Methods("POST").HandlerFunc(PostCreateHandler)

	// Viewing each Post
	post := r.PathPrefix("/posts/{id}/").Subrouter()
	post.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler)
	post.Methods("GET").HandlerFunc(PostShowHandler)
	post.Methods("PUT", "POST").HandlerFunc(PostUpdateHandler)
	post.Methods("DELETE").HandlerFunc(PostDeleteHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Starting server on :" + port)
	http.ListenAndServe(":"+port, r)
}
