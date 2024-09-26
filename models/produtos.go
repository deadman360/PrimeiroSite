package models

import "github.com/deadman360/appWeb/db"

type Produto struct {
	Id         int
	Nome, Desc string
	Preco      float64
	Quantidade int
}

func Search() []Produto {
	p := Produto{}
	produtos := []Produto{
		{Nome: "Camiseta azul", Desc: "Camisa Bonita", Preco: 48.99, Quantidade: 58},
		{Nome: "Camiseta vermelha", Desc: "Camisa Bonita", Preco: 42.99, Quantidade: 598},
		{Nome: "Camiseta verde", Desc: "Camisa Bonita", Preco: 52.99, Quantidade: 13},
	}
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
