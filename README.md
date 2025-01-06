
# todolist-go

Este é um aplicativo de lista de tarefas (To-Do List) desenvolvido em Go (Golang). O aplicativo permite que os usuários criem, visualizem, atualizem e excluam tarefas de forma simples e eficiente.

## Funcionalidades

- Adicionar novas tarefas
- Listar todas as tarefas
- Atualizar tarefas existentes
- Excluir tarefas

## Pré-requisitos

- [Go](https://golang.org/doc/install) instalado
- [Docker](https://docs.docker.com/get-docker/) instalado (opcional, para execução com Docker)

## Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/eduardonakaidev/todolist-go.git
   ```

2. Navegue até o diretório do projeto:

   ```bash
   cd todolist-go
   ```

3. Instale as dependências:

   ```bash
   go mod download
   ```

## Configuração

O aplicativo utiliza um banco de dados SQLite para armazenar as tarefas. Certifique-se de que o arquivo `todolist.db` esteja presente no diretório raiz do projeto.

## Executando o Aplicativo

### Sem Docker

1. Execute o aplicativo:

   ```bash
   go run main.go
   ```

2. O servidor estará disponível em `http://localhost:8080`.

### Com Docker

1. Construa a imagem Docker:

   ```bash
   docker build -t todolist-go .
   ```

2. Inicie o contêiner:

   ```bash
   docker run -p 8080:8080 todolist-go
   ```

3. O servidor estará disponível em `http://localhost:8080`.

## Endpoints da API

- `GET /tasks` - Retorna todas as tarefas
- `GET /tasks/{id}` - Retorna uma tarefa específica por ID
- `POST /tasks` - Cria uma nova tarefa
- `PUT /tasks/{id}` - Atualiza uma tarefa existente
- `DELETE /tasks/{id}` - Exclui uma tarefa

## Estrutura do Projeto

- `controllers/` - Contém os controladores que lidam com as requisições HTTP
- `models/` - Define as estruturas de dados e interações com o banco de dados
- `routes/` - Configura as rotas da aplicação
- `views/` - Contém os templates HTML (se aplicável)
- `main.go` - Ponto de entrada da aplicação

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo `LICENSE` para mais detalhes.
