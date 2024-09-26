package models

import (
	"github.com/deadman360/appWeb/db"
)

type Produto struct {
	Id         int
	Nome, Desc string
	Preco      float64
	Quantidade int
}

func Search() []Produto {
	p := Produto{}
	produtos := []Produto{}

	dbs := db.DbConnect()
	defer dbs.Close()
	query, err := dbs.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = query.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Desc = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}
	return produtos
}

func Create(novo Produto) {
	dbs := db.DbConnect()
	defer dbs.Close()

	insert, err := dbs.Prepare("insert into produtos(nome, descricao, preco,quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(novo.Nome, novo.Desc, novo.Preco, novo.Quantidade)
}

func Delete(idq string) {
	dbs := db.DbConnect()

	delete, err := dbs.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(idq)
}

func Update(idq string) Produto {
	db := db.DbConnect()
	defer db.Close()
	produtoC, err := db.Query("select * from produtos where id=$1", idq)
	if err != nil {
		panic(err.Error())
	}
	var produtoToUp Produto

	for produtoC.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err := produtoC.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoToUp = Produto{Nome: nome, Desc: descricao, Preco: preco, Quantidade: quantidade}
	}
	return produtoToUp

}
