package main

import (
	"bytes"
	// "database/sql"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
	// _ "github.com/lib/pq"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"gopkg.in/unrolled/render.v1"
	"io"
	"mime/multipart"
	"net/http"
	// "os"
	// "reflect"
)

type PostForm struct {
	Photo *multipart.FileHeader
	Name  string
}

func (f *PostForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Photo: "photo",
		&f.Name:  "name",
	}
}

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
	postForm := new(PostForm)
	errs := binding.Bind(r, postForm)
	if errs.Handle(rw) {
		return
	}

	// To access the file data you need to Open the file
	// handler and read the bytes out.
	var fh io.ReadCloser
	var err error
	if fh, err = postForm.Photo.Open(); err != nil {
		http.Error(rw,
			fmt.Sprint("Error opening Mime::Data %+v", err),
			http.StatusInternalServerError)
		return
	}
	defer fh.Close()
	dataBytes := bytes.Buffer{}
	var size int64
	if size, err = dataBytes.ReadFrom(fh); err != nil {
		http.Error(rw,
			fmt.Sprint("Error reading Mime::Data %+v", err),
			http.StatusInternalServerError)
		return
	}

	fileByte := make([]byte, size+100000)
	if _, err = fh.Read(fileByte); err != nil {
		if err != io.EOF {
			http.Error(rw,
				fmt.Sprint("Error reading File %+v", err),
				http.StatusInternalServerError)
			return
		}
	}

	// Now you have the attachment in databytes.
	// Maximum size is default is 10MB.
	fmt.Println("Read %v bytes with filename %s header %s",
		size, postForm.Photo.Filename, postForm.Photo.Header)
	fmt.Fprintln(rw, "photos create")
	fmt.Println(fileByte)
	uploadToS3(fileByte, postForm.Photo.Header.Get("Content-Type"), postForm.Photo.Filename)
}

func uploadToS3(data []byte, contentType string, filename string) {
	auth, err := aws.EnvAuth()
	if err != nil {
		fmt.Println("aws auth error:", err)
	}
	client := s3.New(auth, aws.USWest2)
	bucket := client.Bucket("monoton")
	if err != nil {
		fmt.Println("aws failed to monoton bucket:", err)
	}

	fmt.Println(bucket.Name)

	err = bucket.Put(filename, data, contentType, s3.BucketOwnerFull)
	if err != nil {
		panic(err.Error())
	}
}
