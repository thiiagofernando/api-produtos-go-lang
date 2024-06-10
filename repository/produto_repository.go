package repository

import (
	"api-produtos/constantes"
	"api-produtos/model"
	"database/sql"
	"fmt"
)

type ProdutoRepository struct {
	connection *sql.DB
}

func NewProdutoRepository(connection *sql.DB) ProdutoRepository {
	return ProdutoRepository{
		connection: connection,
	}
}

func (pr *ProdutoRepository) GetProdutos() ([]model.Produto, error) {
	query := "select id,nome,preco from  produto "

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Produto{}, err
	}
	var produtoLista []model.Produto
	var produtoObj model.Produto

	for rows.Next() {
		err = rows.Scan(
			&produtoObj.ID,
			&produtoObj.Nome,
			&produtoObj.Preco)

		if err != nil {
			fmt.Println(err)
			return []model.Produto{}, err
		}

		produtoLista = append(produtoLista, produtoObj)
	}
	rows.Close()
	return produtoLista, nil
}

func (pr *ProdutoRepository) InserirProduto(produto model.Produto) (int64, error) {

	var id int64
	query, err := pr.connection.Prepare("insert into produto" +
		"(nome,preco) values($1, $2) returning id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(produto.Nome, produto.Preco).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()

	return id, nil

}
func (pr *ProdutoRepository) ObterProdutoPorid(id_product int64) (*model.Produto, error) {

	query, err := pr.connection.Prepare("SELECT * FROM produto WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Produto

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Nome,
		&produto.Preco,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}

func (pr *ProdutoRepository) AtualizarProduto(produto model.Produto) (string, error) {

	var id int64
	query, err := pr.connection.Prepare("UPDATE produto" +
		" SET nome= $2, preco= $3 " +
		"WHERE id=$1 ")

	if err != nil {
		fmt.Println(err)
		return constantes.MensagemErroAtualizar, err
	}

	err = query.QueryRow(produto.ID, produto.Nome, produto.Preco).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return constantes.MensagemErroAtualizar, err
	}
	query.Close()
	return constantes.MensagemSucessoAtualizar, err
}

func (pr *ProdutoRepository) ExcluirProdutoPorId(produtoId int64) (string, error) {

	var id int64
	query, err := pr.connection.Prepare("delete from produto where id = $1 ")

	if err != nil {
		fmt.Println(err)
		return constantes.MensagemErroExcluir, err
	}

	err = query.QueryRow(produtoId).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return constantes.MensagemErroExcluir, err
	}
	query.Close()
	return constantes.MensagemSucessoExcluir, err
}
