package repository

import (
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

func (pr *ProdutoRepository) InserirProduto(produto model.Produto) (int, error) {
	
	var id int
	query, err:= pr.connection.Prepare("insert into protudo" + 
						"(nome,preco) values($1, $2) returning id")
	if err != nil {
		fmt.Println(err)
		return 0,err
	}

	err = query.QueryRow(produto.Nome,produto.Preco).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0,err
	}
	query.Close()

	return id,nil

}