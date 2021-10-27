package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login")
	http.HandleFunc("/home")
	http.HandleFunc("/refresh")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
