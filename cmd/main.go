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
	server.GET("/produtos", ProdutoCtrl.ObterProdutos)
	server.POST("/produto", ProdutoCtrl.CriarProduto)
	server.PUT("/produto", ProdutoCtrl.AtualizarProduto)
	server.GET("/produto/:produtoId", ProdutoCtrl.ObterProdutoPorid)
	server.DELETE("/produto/:produtoId", ProdutoCtrl.ExcluirProdutoPorid)

	//Inicia a o servidor http na porta 8080
	server.Run(":8080")
}
