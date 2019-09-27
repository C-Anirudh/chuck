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

func main() {
	r := mux.NewRouter()

	// TODO: Add routing

	log.Fatal(http.ListenAndServe(":9000", r))
}
