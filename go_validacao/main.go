package main

import (
	"api-go-gin-val/database"
	"api-go-gin-val/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
