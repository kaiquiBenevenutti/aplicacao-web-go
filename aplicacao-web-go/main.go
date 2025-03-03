package main

import (
	controller "aplicacao-web-go/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.Index)

	log.Fatal(http.ListenAndServe(":8001", nil))
}
