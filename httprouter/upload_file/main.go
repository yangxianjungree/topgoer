package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	maxUploadSize = 1024 * 1024 * 20 // 20MiB
)

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		log.Printf("File is too big, %v\n", err)
		return
	}

	file, headers, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when we try to get file: %v", err)
		return
	}

	log.Println("file header name: ", headers.Filename, ", size: ", headers.Size, ", header: ", headers.Header)

	// get file's type
	if headers.Header.Get("Content-Type") != "image/png" {
		log.Println("Only upload image file.")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		return
	}

	fn := headers.Filename
	err = ioutil.WriteFile("./video/"+fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully.")
}

func registerHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/upload", uploadHandler)

	return router
}

func main() {
	r := registerHandlers()
	http.ListenAndServe(":8888", r)
}
