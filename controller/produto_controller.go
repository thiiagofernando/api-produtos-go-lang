package controller

import (
	"api-produtos/model"
	"api-produtos/usecase"
	"net/http"
	"strconv"

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

func (p *produtoController) ObterProdutos(ctx *gin.Context) {

	produtos, err := p.produtoUseCase.GetProdutos()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, produtos)
}

func (p *produtoController) CriarProduto(ctx *gin.Context) {

	var product model.Produto
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.produtoUseCase.CriarProduto(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *produtoController) ObterProdutoPorid(ctx *gin.Context) {

	id := ctx.Param("produtoId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.produtoUseCase.ObterProdutoPorid(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *produtoController) AtualizarProduto(ctx *gin.Context) {

	var product model.Produto
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	mensagem, err := p.produtoUseCase.AtualizarProduto(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, mensagem)
}

func (p *produtoController) ExcluirProdutoPorid(ctx *gin.Context) {

	id := ctx.Param("produtoId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	mensagem, err := p.produtoUseCase.ExcluirProdutoPorid(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, mensagem)
}
