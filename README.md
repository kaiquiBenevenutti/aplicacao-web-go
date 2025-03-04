# Aplicação Web em Go - Controle de Produtos

**aplicacao-web-go** é um projeto desenvolvido em Go para fazer o controle de produtos, podendo criar novos produtos, editar produtos já cadastrados e deletar.

## Funcionalidades

- Criação de produtos com insert no banco de dados.
- Lista com todos os produtos retornados pelo banco em forma de tabela e ordenados por id.
- Edição de produtos podendo trocar informações separadamente.

## Como funciona

O site possui 3 telas:

1. **Pagina principal**: apresenta uma tabela com os produtos do banco de dados com botoes de atualizar pagina, criar novo produto, editar e deletar.
2. **Criar produto**: Mostra um formulario com campos para criação de produtos.
3. **Editar produto**: Uma pagina semelhante a de criação, mas as informações do produto selecionado são preenchidas no formulario logo ao entrar na pagina, podendo alterar alguma informação e salvar.

## Requisitos

- [Go](https://golang.org/dl/) instalado na sua máquina.
- Banco de dados PostgreSQL configurado com o localhost .

## Como executar

1. Clone o repositório:
   
   ```bash
   git clone https://github.com/kaiquiBenevenutti/aplicacao-web-go.git

2. Navegue até o diretório do projeto:
   ```bash
   cd aplicacao-web-go

3. Rode o programa usando o comando:
   ```bash
   go run main.go

4. Entre no seu navegador e coloque a porta: http://localhost:8001/

### Estrutura do Projeto

 ```bash
aplicacao-web-go/
├── main.go             # Arquivo principal da aplicação
├── controllers/        # Controladores para gerenciar as requisições HTTP
├── models/             # Definição dos modelos (ex.: Produto)
├── routes/             # Rotas da aplicação
├── db/                 # Configurações, como conexão com o banco de dados
└── templates/          # Paginas HTML

```
### Autor

- Kaiqui Benevenutti
