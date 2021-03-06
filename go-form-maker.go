package main

import (
	"encoding/json"
	"fmt"
	"github.com/Maksadbek/go-form-maker/form"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting...")
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8800", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var myform form.MyForm = form.MyForm{Age: 18, Token: "deadbeef"}
	formxml, err := form.FormCreate(&myform)
	if err != nil {
		log.Println("error while building form")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, formxml)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	var myform form.MyForm
	err := form.FormRead(&myform, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jform, err := json.MarshalIndent(myform, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(string(jform))
	fmt.Fprintf(w, string(jform))
}
