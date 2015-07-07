package main

import (
	// "database/sql"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	// _ "github.com/lib/pq"
	// "github.com/mitchellh/goamz/aws"
	// "github.com/mitchellh/goamz/s3"
	"gopkg.in/unrolled/render.v1"
	"net/http"
	// "os"
)

var renderer = render.New(render.Options{})

func main() {
	fmt.Println("starting monoton")

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", HomeHandler)

	posts := r.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(PostsIndexHandler)

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir(".")),
	)
	n.UseHandler(r)

	n.Run(":3000")
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	renderer.HTML(rw, http.StatusOK, "index", nil)
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts index")
}
