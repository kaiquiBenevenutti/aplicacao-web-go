package models

import (
	db2 "aplicacao-web-go/db"
	"database/sql"
	"log"
	"time"
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

	selectProducts, err := db.Query("SELECT * FROM produtos where exclusao is null ")

	if err != nil {
		log.Fatal(err)
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var valor float64
		var exclusao sql.NullTime

		err := selectProducts.Scan(&id, &nome, &descricao, &valor, &quantidade, &exclusao)
		if err != nil {
			log.Fatal(err)
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Valor = valor
		p.Quantidade = quantidade
		p.Id = id

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

func DeleteProduto(id string) {
	db := db2.ConectionDB()
	executa, err := db.Prepare("UPDATE produtos SET exclusao = $2 WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}
	executa.Exec(id, time.Now())
	defer db.Close()
}
