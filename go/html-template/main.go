package main

import (
	"html/template"
	"net/http"
	"os"
	"path"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func foo(rw http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(rw, profile); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", foo)
	http.ListenAndServe(":"+port, nil)
}
