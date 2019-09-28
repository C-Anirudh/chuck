package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, data interface{}) error {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	dec := schema.NewDecoder()
	if err := dec.Decode(data, r.PostForm); err != nil {
		panic(err)
	}
	return nil
}
