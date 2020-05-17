package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", GreetingServer)
	http.ListenAndServe(":8000", nil)
}

func GreetingServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}