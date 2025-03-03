package models

import (
	"github.com/kaiquiBenevenutti/db"
	"log"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

func getProdutos() []Produto {
	db := ConectionDB()

	selectProducts, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		log.Fatal(err)
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var valor float64

		err := selectProducts.Scan(&id, &nome, &descricao, &valor, &quantidade)
		if err != nil {
			log.Fatal(err)
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Valor = valor
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
