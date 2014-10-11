package main

import (
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

func main() {
	r := render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, re *http.Request) {
		rw.Write([]byte("Welcome, visit sub pages now."))
	})

	mux.HandleFunc("/data", func(rw http.ResponseWriter, re *http.Request) {
		r.Data(rw, http.StatusOK, []byte("Some binary data here."))
	})

	mux.HandleFunc("/json", func(rw http.ResponseWriter, re *http.Request) {
		r.JSON(rw, http.StatusOK, map[string]string{"hello": "json"})
	})

	mux.HandleFunc("/html", func(rw http.ResponseWriter, re *http.Request) {
		r.HTML(rw, http.StatusOK, "example", nil)
	})

	http.ListenAndServe(":3000", mux)
}
