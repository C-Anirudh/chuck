package main

import (
	"log"
	"net/http"

	"github.com/C-Anirudh/chuck/service/contollers"
	"github.com/gorilla/mux"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func main() {
	r := mux.NewRouter()

	// TODO: Add routing
	r.HandleFunc("/login", contollers.LoginHandler).Methods("POST")
	r.HandleFunc("/signup", contollers.SignupHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}
