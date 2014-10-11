package main

import "net/http"

// Runs on address we give it, and when it recieves on HTTP request it will and it off to the http.Handler that we supply.
func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
