package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/C-Anirudh/chuck/service/controllers"
	"github.com/C-Anirudh/chuck/service/models"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = ""
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()

	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()

	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/signup", usersC.Signup).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}
