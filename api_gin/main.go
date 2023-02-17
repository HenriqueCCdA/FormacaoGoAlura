package main

import (
	"api-rest-gin/models"
	"api-rest-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Gui Lima", CPF: "32132132111", RG: "47000000000"},
		{Nome: "Ana", CPF: "12312312322", RG: "48000000000"},
	}
	routes.HandleRequests()
}
