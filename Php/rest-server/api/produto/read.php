<?php
    header('Access-Control-Allow-Origin: *');
    header('Content-Type: application/json');

    include_once '../../config/Database.php';
    include_once '../../models/Produto.php';

    $database = new Database();
    $db = $database->conectar();

    $produto = new Produto($db);

    $rs = $produto->read();

    $num = $rs->rowCount();

    if($num > 0) {
        
        $produtos = array();
        $produtos['data'] = array();

        while($row = $rs->fetch(PDO::FETCH_ASSOC)) {
            extract($row);

            $p_item = array(
                'id'         => $id,
                'descricao'  => $descricao,
                'quantidade' => $quantidade,
                'valor'      => $valor
            );

            array_push($produtos['data'], $p_item);
        }
        echo json_encode($produtos, JSON_PRETTY_PRINT);
    } else {
        echo json_encode (
            array('mensagem' => "Produtos not found")
        );
    }

    /* 
        Como usar:
        
        http://localhost/rest-server/api/produto/read
    
    */