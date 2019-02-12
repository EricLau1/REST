Blockcoin - Go Api Restful - Para Estudo
=========================================

Endpoints:

    GET/  http://localhost:3000/
    |
    |_ Home page

    POST/ http://localhost:3000/api/users
    |
    |_ Recebe um json com nickname, email e senha.

    POST/ http://localhost:3000/apí/login
    |
    |_ Recebe um json com email e senha para autenticação
    |_ Retorna um JWT

    GET/ http://localhost:3000/api/users
    |
    |_ Retorna um json com uma lista de usuários

    POST/ http://localhost:3000/api/users
    |
    |_ Recebe um json com nickname, email e senha
    |_ Precisa de Autenticação

    GET/ http://localhost:3000/api/users/1
    |
    |_ Retorna um usuário buscando pelo Id
    |_ Precisa de Autenticação

    PUT/ http://localhost:3000/api/users/1
    |
    |_ Recebe um json com os dados que serão atualizados do usuário que possui o Id informado na Url
    |_ Retorna a quantidade de linhas atualizadas
    |_ Precisa de Autenticação

    DELETE/ http://localhost:3000/api/users/1
    |
    |_ Deleta um usuário informando Id pela Url
    |_ Retorna quantidade de linhas afetadas
    |_ Precisa de Autenticação

    GET/ http://localhost:3000/api/wallets
    |
    |_ Retorna uma lista de carteiras

    GET/ http://localhost:3000/api/wallets/public_key
    |
    |_ Retorna uma carteira informando a chave publica na Url

    PUT/ http://localhost:3000/api/wallets/public_key
    |
    |_ Atualiza o saldo de uma carteira informando a chave publica na Url

    PUT/ http://localhost:3000/api/wallets/add/public_key
    |
    |_ Adiciona um valor via json ao saldo da carteira informando a chave publica na Url
    |_ Precisa de autenticação

    GET/ http://localhost:3000/api/transactions
    |
    |_ Retorna um json com uma lista de transações

    POST/ http://localhost:3000/api/transactions/public_key
    |
    |_ Recebe um json com a chave pública e um valor de uma carteira que será
    |  trasferido para a carteira cuja chave pública e passada pela url
    |_ Precisa de autenticação