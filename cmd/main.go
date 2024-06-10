package main

import (
	"api-produtos/controller"
	"api-produtos/db"
	"api-produtos/repository"
	"api-produtos/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConectDB()
	if err != nil {
		panic(err)
	}

	ProductRepo := repository.NewProdutoRepository(dbConnection)
	ProdutoUseCase := usecase.NewProdutoUseCase(ProductRepo)
	ProdutoCtrl := controller.NewProdutoController(ProdutoUseCase)

	//Rotas API
	server.GET("produto", ProdutoCtrl.GetProdutos)

	//Inicia a o servidor http na porta 8080
	server.Run(":8080")
}
