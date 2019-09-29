package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/C-Anirudh/chuck/service/models"
)

// NewUsers parses the templates related to the user and stores them in Users struct
func NewUsers(us models.UserService) *Users {
	return &Users{
		us: us,
	}
}

// Login handles the login form
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("In login handler")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid Email address")
		case models.ErrPasswordIncorrect:
			fmt.Fprintln(w, "Invalid Password Provided")
		case nil:
			fmt.Fprintln(w, user)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Write([]byte("true"))
}

// Signup will parse the sign up form and create a new user
func (u *Users) Signup(w http.ResponseWriter, r *http.Request) {
	log.Println("In signup handler")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		fmt.Println(w, err)
		return
	}
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := u.us.Create(&user); err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("true"))

}

// Users will hold the user service
type Users struct {
	us models.UserService
}

// SignupForm contains the details entered by the user in the signup form
type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// LoginForm contains the details entered by the user in the login form
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
