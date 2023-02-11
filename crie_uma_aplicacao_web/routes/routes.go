package routes

import (
	"net/http"
	"web/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
