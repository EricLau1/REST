<?php
    header('Access-Control-Allow-Origin: *');
    header('Content-Type: application/json');
    header('Access-Control-Allow-Methods: DELETE');
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

    $produto->setId($data->id);

    if($produto->delete()) {
        echo json_encode(
            array("mensagem" => "Produto removido com sucesso!")
        );
    } else {
        echo json_encode(
            array("mensagem" => "Produto nao foi removido.")
        );
    }

    /* 
    
        Como usar:

        Baixe o Postman : https://www.getpostman.com/

        Procure pelo ID do item que deseja excluir

        Abra o programa

        Em Headers na opção Key selecione o item Content-Type,
        e na opção Value selecione o application/json

        Mude a opção na URL para PUT

        E cole a URL abaixo:

        http://localhost/rest-server/api/produto/delete

        Clique em Body e selecione a opção RAW

        Coleque a estrutura do objeto em formato json:

{
    "id" : 13
}


        Clique em SEND
    */   