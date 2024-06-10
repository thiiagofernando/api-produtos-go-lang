package usecase

import (
	"api-produtos/model"
	"api-produtos/repository"
)

type ProdutoUseCase struct {
	repository repository.ProdutoRepository
}

func NewProdutoUseCase(repo repository.ProdutoRepository) ProdutoUseCase {
	return ProdutoUseCase{
		repository: repo,
	}
}

func (pu *ProdutoUseCase) GetProdutos() ([]model.Produto, error) {
	return pu.repository.GetProdutos()
}
