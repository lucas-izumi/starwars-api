# starwars-api

Este projeto mantém uma base de dados de planetas do universo de Star Wars, permitindo busca, inserção e deleção dos registros. A persistência dos registros é por sessão.
O projeto foi implementado como uma API REST usando a linguagem Go.

# Exemplo de uso
## Inserção
- Utilizar o endpoint `/planets/` com o comando **POST** passando um JSON com os dados `name`, `climate` e `terrain`. O ID e número de aparições nos filmes são preenchidos automaticamente.

## Busca
- Utilizar o endpoint `/planets/` com o comando **GET** passando os parâmetros da busca: nome do planeta ou ID. Exemplo: `/planets/1` busca o planeta de ID 1. `/planets/tatooine` busca o planeta de nome Tatooine.
- Para buscar todos os registros, basta usar o endpoint `/planets` com um comando **GET**

## Deleção
- Utilizar o endpoint `/planets/` com o comando **DEL** passando o ID do planeta a ser apagado. Exemplo: `/planets/1` deleta o planeta de ID 1.
