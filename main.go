package main

import (
	"fmt"
	"net/http"
	"gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Query().Get("id"))
}

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/getrequest", GetRequest)

	http.HandleFunc("/postrequest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello, you've requested: %s\n", r.FormValue("email"), r.FormValue("Username"))
		return "ok"
	})

	http.ListenAndServe(":8080", nil)
}
