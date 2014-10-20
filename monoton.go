package main

import (
	// "database/sql"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	// _ "github.com/lib/pq"
	"net/http"
	"os"
)

func main() {
	fmt.Println("starting monoton")
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/login", UserLoginHandler).Methods("POST")
	r.HandleFunc("/logout", UserLogoutHandler).Methods("POST")
	r.HandleFunc("/signup", UserSignupHandler).Methods("POST")
	r.HandleFunc("/showme", UserShowmeHandler).Methods("GET")

	posts := r.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(PostsIndexHandler)
	posts.Methods("POST").HandlerFunc(PostsCreateHandler)
	posts.Methods("DELETE").HandlerFunc(PostsDeleteHandler)

	n := negroni.Classic()
	n.UseHandler(r)

	n.Run(":" + os.Getenv("PORT"))
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Monoton")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts index")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts create")
}

func PostsDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts delete")
}

func UserLoginHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "user login")
}

func UserLogoutHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "user logout")
}

func UserSignupHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "user signup")
}

func UserShowmeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "user show me")
}
