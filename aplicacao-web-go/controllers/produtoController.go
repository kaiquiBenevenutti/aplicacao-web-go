package controllers

import (
	model "aplicacao-web-go/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := model.GetProdutos()
	temp.ExecuteTemplate(w, "index", produtos)
}
