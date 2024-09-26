package routes

import (
	"net/http"

	"github.com/deadman360/appWeb/controller"
)

func Router() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/update", controller.Update)
}
