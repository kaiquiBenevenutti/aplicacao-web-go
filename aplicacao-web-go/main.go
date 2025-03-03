package main

import (
	"aplicacao-web-go/routes"
	"log"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	log.Fatal(http.ListenAndServe(":8001", nil))
}
