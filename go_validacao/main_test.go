package main

import (
	"api-go-gin-val/controllers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, http.StatusOK, resp.Code, "Deveriam ser iguais")
	mockDaReposta := `{"API diz:":"E ai gui, tudo beleza?"}`
	repostaBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockDaReposta, string(repostaBody))

}
