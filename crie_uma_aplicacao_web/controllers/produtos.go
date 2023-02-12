package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão do quantidae:", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)

	}
	http.Redirect(w, r, "Index", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idDoProduto)

	http.Redirect(w, r, "Index", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id")

	produto := models.EditaProduto(idDoProduto)
	fmt.Println(produto)

	temp.ExecuteTemplate(w, "Edit", produto)
}