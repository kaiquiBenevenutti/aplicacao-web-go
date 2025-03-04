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

	selectProducts, err := db.Query("SELECT * FROM produtos where exclusao is null order by id")

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

func EditaProduto(idProduto string) Produto {
	db := db2.ConectionDB()
	produtoBanco, _ := db.Query("SELECT * FROM produtos where exclusao is null and id=$1", idProduto)
	p := Produto{}
	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var valor float64
		var exclusao sql.NullTime

		produtoBanco.Scan(&id, &nome, &descricao, &valor, &quantidade, &exclusao)

		p.Nome = nome
		p.Descricao = descricao
		p.Valor = valor
		p.Quantidade = quantidade
		p.Id = id
	}
	defer db.Close()
	return p
}

func UpdateProduto(produto Produto) {
	db := db2.ConectionDB()

	concection, _ := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	concection.Exec(produto.Nome, produto.Descricao, produto.Valor, produto.Quantidade, produto.Id)

	defer db.Close()
}
