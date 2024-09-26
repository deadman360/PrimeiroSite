package main

import (
	"fmt"
	"net/http"

	"github.com/deadman360/appWeb/routes"
	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Server rodando!!")
	routes.Router()
	http.ListenAndServe(":8000", nil)

}
