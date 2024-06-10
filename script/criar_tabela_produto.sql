CREATE TABLE IF NOT EXISTS produto (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(255) NOT null,
        preco NUMERIC(10,2) NOT null
        );