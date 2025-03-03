package controllers

import (
	model "aplicacao-web-go/models"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := model.GetProdutos()
	temp.ExecuteTemplate(w, "index", produtos)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new-produto", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var produto model.Produto
		produto.Nome = r.FormValue("nome")
		Valor := r.FormValue("preco")
		produto.Descricao = r.FormValue("descricao")
		Quantidade := r.FormValue("quantidade")

		produto.Valor, _ = strconv.ParseFloat(Valor, 64)
		produto.Quantidade, _ = strconv.Atoi(Quantidade)

		model.PostProdutos(produto)
	}
	http.Redirect(w, r, "/", 301)
}
