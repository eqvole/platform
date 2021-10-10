package handlers

import "net/http"

type User struct{}

func NewUserHandler() User {
	return User{}
}

func (u *User) Register(w http.ResponseWriter, r *http.Request) {

}

func (u *User) Auth(w http.ResponseWriter, r *http.Request) {

}
