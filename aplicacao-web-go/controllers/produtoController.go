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

func Delete(w http.ResponseWriter, r *http.Request) {
	produtoDeleteId := r.URL.Query().Get("id")
	model.DeleteProduto(produtoDeleteId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := model.EditaProduto(idProduto)
	temp.ExecuteTemplate(w, "edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var produto model.Produto
		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")
		Quantidade := r.FormValue("quantidade")
		Valor := r.FormValue("preco")
		id := r.FormValue("id")

		produto.Valor, _ = strconv.ParseFloat(Valor, 64)
		produto.Quantidade, _ = strconv.Atoi(Quantidade)
		produto.Id, _ = strconv.Atoi(id)

		model.UpdateProduto(produto)
	}
	http.Redirect(w, r, "/", 301)
}
