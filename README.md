# 📦 API REST em Go com Gin e MySQL

Este projeto é uma API REST desenvolvida em **Go** utilizando o framework **Gin** e o ORM **GORM** para interagir com um banco de dados **MySQL**. A API permite realizar operações **CRUD (Create, Read, Update, Delete)** em produtos.

## 🚀 Tecnologias Utilizadas
- **Go** (Golang)
- **Gin** (Framework Web)
- **GORM** (ORM para Go)
- **MySQL** (Banco de Dados)
- **HeidiSQL** (Cliente para gerenciar o banco)

## 📌 Funcionalidades
- Criar um novo produto
- Listar todos os produtos
- Buscar um produto por ID
- Atualizar um produto existente
- Deletar um produto

## 🛠 Configuração e Execução

### 1️⃣ Clone o repositório
```sh
git clone https://github.com/seu-usuario/project_api_go.git
cd project_api_go
```

### 2️⃣ Configure o banco de dados MySQL
Crie um banco de dados e ajuste o **DSN** no arquivo `database/database.go`:
```go
dsn := "usuario:senha@tcp(127.0.0.1:3306)/nome_do_banco?charset=utf8mb4&parseTime=True&loc=Local"
```

### 3️⃣ Instale as dependências
```sh
go mod tidy
```

### 4️⃣ Execute a API
```sh
go run main.go
```

### 5️⃣ Testando a API
Você pode usar o **Postman**, **Insomnia** ou **cURL** para testar os endpoints.

## 🔗 Endpoints da API

| Método  | Endpoint        | Descrição                 |
|---------|---------------|---------------------------|
| `GET`   | `/products`   | Listar todos os produtos |
| `GET`   | `/products/:id` | Buscar um produto por ID |
| `POST`  | `/products`   | Criar um novo produto |
| `PUT`   | `/products/:id` | Atualizar um produto |
| `DELETE` | `/products/:id` | Remover um produto |



