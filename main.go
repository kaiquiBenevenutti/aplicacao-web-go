package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

func conectionDB() *sql.DB {
	connStr := "user=postgres dbname=web_aplicacao_go port=5432 sslmode=disable password=Kaiquirb1!"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func index(w http.ResponseWriter, r *http.Request) {

	db := conectionDB()

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

	temp.ExecuteTemplate(w, "index", produtos)
	defer db.Close()
}
