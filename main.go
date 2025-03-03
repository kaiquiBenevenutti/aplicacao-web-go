package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func index(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "index", produtos)
	defer db.Close()
}
