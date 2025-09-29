<p align="center">
  <h1 align="center">🛍️ Products API GO</h1>
  <p align="center">
    Uma API RESTful para gerenciamento de produtos, desenvolvida em Go.
  </p>
</p>

---

<p align="center">
  <a href="#descrição">Descrição</a> |
  <a href="#tecnologias-utilizadas">Tecnologias Utilizadas</a> |
  <a href="#como-rodar-e-testar-localmente">Como Rodar e Testar Localmente</a> |
  <a href="#documentação-das-rotas">Documentação das Rotas</a>
</p>

---

## 📝 Descrição

A **Products API GO** é uma API RESTful desenvolvida em GO como parte dos meus estudos e aprimoramento profissional. O projeto tem como objetivo testar minhas habilidades desenvolvendo uma aplicação (CRUD) utilizando boas práticas da linguagem GO

---

## 🚀 Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gin](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)

---

## 🐳 Como Rodar e Testar Localmente

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu-usuario/products-api-go.git
   cd products-api-go
   ```

2. **Suba os containers com Docker Compose:**

   ```bash
   docker-compose up -d
   ```

3. **Acesse a API:**

   - A API estará disponível em `http://localhost:8080`

---

## 📚 Documentação das Rotas

### Produtos

- **GET `/products`**

  Lista todos os produtos.

  - **Exemplo de resposta:**
    ```json
    [
      {
        "id": 1,
        "name": "Teclado Mecânico Keychron K2",
        "description": "Teclado mecânico compacto com Bluetooth",
        "price": 399.9,
        "slug": "teclado-mecanico-keychron-k2"
      },
      {
        "id": 2,
        "name": "Mouse Logitech MX Master 3",
        "description": "Mouse sem fio avançado com rolagem eletromagnética",
        "price": 499.9,
        "slug": "mouse-logitech-mx-master-3"
      }
      ...
    ]
    ```

- **GET `/products/{id}`**

  Busca um produto pelo ID.

  - **Exemplo de resposta:**
    ```json
    {
      "id": 2,
      "name": "Mouse Logitech MX Master 3",
      "description": "Mouse sem fio avançado com rolagem eletromagnética",
      "price": 499.9,
      "slug": "mouse-logitech-mx-master-3"
    }
    ```

- **POST `/products`**

  Cria um novo produto.

  - **Body exemplo:**
    ```json
    {
      "name": "Produto Exemplo",
      "description": "Descrição do produto exemplo",
      "price": 99.9,
      "slug": "produto-exemplo"
    }
    ```

- **PUT `/products/{id}`**

  🚧 Está sendo desenvolvido...

- **DELETE `/products/{id}`**

  🚧 Está sendo desenvolvido...

---

<p align="center">
  Feito com 💙 por
  <a href="https://www.linkedin.com/in/goncadanilo" target="_blank">Danilo Santos</a>!
</p>
