<?php
    header('Access-Control-Allow-Origin: *');
    header('Content-Type: application/json');
    header('Access-Control-Allow-Methods: PUT');
    header('Access-Control-Allow-Headers: 
    Access-Control-Allow-Headers, 
    Content-Type, 
    Access-Control-Allow-Methods,
    Authorization, X-Requested-With');

    include_once '../../config/Database.php';
    include_once '../../models/Produto.php';

    $database = new Database();
    $db = $database->conectar();

    $produto = new Produto($db);

    $data = json_decode(file_get_contents('php://input'));

    $produto->setDescricao($data->descricao);
    $produto->setQuantidade($data->quantidade);
    $produto->setValor($data->valor);
    $produto->setId($data->id);

    if($produto->update()) {
        echo json_encode(
            array("mensagem" => "Produto modificado com sucesso!")
        );
    } else {
        echo json_encode(
            array("mensagem" => "Produto nao foi modificado.")
        );
    }

    /* 
    
        Como usar:

        Baixe o Postman : https://www.getpostman.com/

        Procure pelo ID do item que deseja modificar

        Abra o programa

        Em Headers na opção Key selecione o item Content-Type,
        e na opção Value selecione o application/json

        Mude a opção na URL para PUT

        E cole a URL abaixo:

        http://localhost/rest-server/api/produto/update

        Clique em Body e selecione a opção RAW

        Coleque a estrutura do objeto em formato json:

{
    "id" : 13,
    "descricao" : "item-teste",
    "quantidade" : 20,
    "valor" : 100
}


        Clique em SEND
    */   