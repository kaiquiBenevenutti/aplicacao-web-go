package main

import (
	"html/template"
	"log"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Cachorro quente", "Muito saboroso", 15.00, 25},
		{"Camiseta", "Bem bonita", 29.99, 10},
	}

	err := temp.ExecuteTemplate(w, "index", produtos)
	if err != nil {
		http.Error(w, "Erro ao renderizar a página", http.StatusInternalServerError)
		log.Println("Erro no template:", err)
	}
}
