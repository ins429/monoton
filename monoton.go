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

	photos := r.Path("/photos").Subrouter()
	photos.Methods("GET").HandlerFunc(PhotosIndexHandler)
	photos.Methods("POST").HandlerFunc(PhotosCreateHandler)

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

func PhotosIndexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "photos index")
}

func PhotosCreateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "photos create")
}
