package main

import (
	"api-go-gin-val/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTeste() *gin.Engine {

	rotas := gin.Default()

	return rotas

}

func TestStatusCodeDaSaudacaoComParamentro(t *testing.T) {
	r := SetupDasRotasDeTeste()

	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/gui", nil)

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Status error: valor  recebido foi %d e o esperado era %d", resp.Code, http.StatusOK)
	}
}
