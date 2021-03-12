package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	fmt.Println(r.PostForm)

	decoder := schema.NewDecoder()
	if err := decoder.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
