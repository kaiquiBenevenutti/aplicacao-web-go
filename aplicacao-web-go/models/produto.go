package models

import (
	db2 "aplicacao-web-go/db"
	"log"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

func GetProdutos() []Produto {
	db := db2.ConectionDB()

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

func PostProdutos(produto Produto) {
	db := db2.ConectionDB()
	executa, err := db.Prepare("INSERT INTO PRODUTOS(nome,  descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	executa.Exec(produto.Nome, produto.Descricao, produto.Valor, produto.Quantidade)
	defer db.Close()
}
