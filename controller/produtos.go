package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/deadman360/appWeb/models"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.Search()
	templ.ExecuteTemplate(w, "Index", produtos)

}
func New(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "New", nil)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err.Error())
		}
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão da descrição:", err.Error())
		}
		novoProduto := models.Produto{Nome: nome, Desc: descricao, Preco: preco, Quantidade: quantidade}
		models.Create(novoProduto)
	}
	http.Redirect(w, r, "/", 301)
}
