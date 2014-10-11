package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func ProfileHandler(rw http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", ProfileHandler)
	http.ListenAndServe(":"+port, nil)
}
