package routes

import (
	controller "aplicacao-web-go/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new-produto", controller.NewProduct)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.Update)
}
