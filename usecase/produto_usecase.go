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

func (pu *ProdutoUseCase) CriarProduto(product model.Produto) (model.Produto, error) {

	productId, err := pu.repository.InserirProduto(product)
	if err != nil {
		return model.Produto{}, err
	}

	product.ID = productId

	return product, nil
}
func (pu *ProdutoUseCase) ObterProdutoPorid(id_product int64) (*model.Produto, error) {

	product, err := pu.repository.ObterProdutoPorid(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProdutoUseCase) AtualizarProduto(produto model.Produto) bool {

	retorno := pu.repository.AtualizarProduto(produto.ID, produto.Nome, produto.Preco)
	return retorno
}
func (pu *ProdutoUseCase) ExcluirProduto(id_product int64) bool {

	retorno := pu.repository.ExcluirProduto(id_product)
	return retorno
}
