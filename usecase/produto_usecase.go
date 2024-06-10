package usecase

import (
	"api-produtos/constantes"
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

func (pu *ProdutoUseCase) AtualizarProduto(produto model.Produto) (string, error) {

	var mensagem string
	mensagem, err := pu.repository.AtualizarProduto(produto)
	if err != nil {
		return constantes.MensagemErroAtualizar, err
	}

	return mensagem, nil
}
func (pu *ProdutoUseCase) ExcluirProdutoPorid(id_product int64) (string, error) {

	mensagem, err := pu.repository.ExcluirProdutoPorId(id_product)
	if err != nil {
		return constantes.MensagemErroExcluir, err
	}

	return mensagem, nil
}
