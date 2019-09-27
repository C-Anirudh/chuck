package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In login handler")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte("true"))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In signup handler")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte("true"))
}

func main() {
	r := mux.NewRouter()

	// TODO: Add routing
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/signup", signupHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}
