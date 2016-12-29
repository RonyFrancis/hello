package app

import (
	"controller"
	"net/http"
)

func Run() {
	http.HandleFunc("/view/", controller.ViewHandler)
	http.HandleFunc("/edit/", controller.EditHandler)
	http.HandleFunc("/save/", controller.SaveHandler)
}
