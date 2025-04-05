# 💸 Desafio Go: Cotação do Dólar (Client/Server com Context, HTTP e SQLite)

Este projeto foi desenvolvido como parte de um desafio prático com o objetivo de aplicar conceitos fundamentais da linguagem Go, incluindo:

- Web server com HTTP
- Consumo de API externa
- Manipulação de arquivos
- Persistência em banco de dados (SQLite)
- Gerenciamento de tempo com `context`

## 📄 Descrição

O sistema consiste em dois programas:

- `server.go`: Servidor HTTP que expõe um endpoint `/cotacao`. Ao receber uma requisição, ele consulta a cotação atual do dólar em uma API externa e grava essa cotação no banco de dados SQLite.
- `client.go`: Cliente HTTP que requisita a cotação ao servidor, exibe o valor retornado e salva a cotação no arquivo `cotacao.txt`.

## 📦 Requisitos

- Go 1.18+
- SQLite (ou a lib embutida no Go via `github.com/mattn/go-sqlite3`)
- Conexão com a internet (para acessar a API de câmbio)

## 🚀 Como executar

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/nome-do-repo.git
cd nome-do-repo
```
### 2. Instale as dependências
bash
Copiar
Editar
go mod tidy

### 3. Execute o servidor
bash
Copiar
Editar
go run server.go
O servidor ficará disponível em http://localhost:8080.

### 4. Em outro terminal, execute o cliente
bash
Copiar
Editar
go run client.go
O cliente irá:

Fazer uma requisição para http://localhost:8080/cotacao

Salvar o valor retornado no arquivo cotacao.txt no formato:

makefile
Copiar
Editar
Dólar: 5.23
🧪 Testando com Postman
Você pode testar o servidor diretamente com o Postman:

Método: GET

URL: http://localhost:8080/cotacao

Resposta esperada:

json
Copiar
Editar
{
  "bid": "5.23"
}
### 🕐 Timeouts implementados
server.go:

Timeout de 200ms para chamada da API externa

Timeout de 10ms para gravação no banco

client.go:

Timeout de 300ms para resposta do servidor

Todos os contextos geram logs de erro se os limites forem ultrapassados.

### 🗃 Estrutura do banco
O SQLite cria automaticamente a tabela cotacoes com os seguintes campos:

sql
Copiar
Editar
id INTEGER PRIMARY KEY
bid TEXT
created_at DATETIME DEFAULT CURRENT_TIMESTAMP
📁 Arquivo gerado
O cliente salva o arquivo cotacao.txt com o conteúdo da cotação mais recente no formato:

css
Copiar
Editar
Dólar: {valor}
📌 Observações
O projeto utiliza apenas bibliotecas padrão do Go (exceto o driver SQLite).

Todos os erros de contexto e operações críticas são logados no terminal.

A API utilizada para cotação é pública: AwesomeAPI

### 🧑‍💻 Autor
Desenvolvido por Kaue P. da Silva
