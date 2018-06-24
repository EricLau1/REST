<?php
    header('Access-Control-Allow-Origin: *');
    header('Content-Type: application/json');

    include_once '../../config/Database.php';
    include_once '../../models/Produto.php';

    $database = new Database();
    $db = $database->conectar();

    $produto = new Produto($db);

    $id = isset($_GET['id']) ? $_GET['id'] : die();
    
    $produto->setId($id);
    
    $produto->findOne();

    $p_item = array(
        'id'         => $produto->getId(),
        'descricao'  => $produto->getDescricao(),
        'quantidade' => $produto->getQuantidade(),
        'valor'      => $produto->getValor()
    );

    print_r(json_encode($p_item, JSON_PRETTY_PRINT));

    /* 
        Como usar:
        
        http://localhost/rest-server/api/produto/findOne?id=1
    
    */