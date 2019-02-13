package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/myname/", myName)
	// http.HandleFunc("/whatsyourname", talha)
	http.Handle("/myname", http.HandlerFunc(myName))
	http.Handle("/whatsyourname", http.HandlerFunc(talha))
	// http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func myName(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Muhammad Talha Umar")
}

func talha(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "something.gohtml", "My Name is Talha")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
