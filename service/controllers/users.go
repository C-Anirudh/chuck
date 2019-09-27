package controllers

import (
	"log"
	"net/http"
)

// LoginHandler handles the login form
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In login handler")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte("true"))
}

// SignupHandler registers a new user according to the signup form filled out by the user
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In signup handler")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte("true"))
}
