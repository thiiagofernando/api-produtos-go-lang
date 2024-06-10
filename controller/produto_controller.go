package controller

import (
	"api-produtos/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type produtoController struct {
	produtoUseCase usecase.ProdutoUseCase
}

func NewProdutoController(useCase usecase.ProdutoUseCase) produtoController {
	return produtoController{
		produtoUseCase: useCase,
	}
}

func (p *produtoController) GetProdutos(ctx *gin.Context) {

	produtos, err := p.produtoUseCase.GetProdutos()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, produtos)
}
