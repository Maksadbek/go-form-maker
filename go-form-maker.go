package main

import (
	"fmt"
	"github.com/Maksadbek/go-form-maker/form"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting...")
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8800", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var myform form.MyForm
	formxml, err := form.FormCreate(&myform)
	if err != nil {
		log.Println("error while building form")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println(formxml)
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, formxml)
}
