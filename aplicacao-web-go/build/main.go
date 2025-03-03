package main

import (
	model "aplicacao-web-go/models"
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
	produtos := model.GetProdutos()
	temp.ExecuteTemplate(w, "index", produtos)
}
