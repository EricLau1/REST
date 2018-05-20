<?php

$descricao = addslashes(trim($_POST['descricao']));
$quantidade = addslashes(trim($_POST['quantidade']));
$valor = addslashes(trim($_POST['valor']));

echo "$descricao, $quantidade, $valor";

if(!empty($descricao) && !empty($quantidade) && !empty($valor)) {

  $url = "http://localhost:3000/produtos";

  $cliente = curl_init($url);

  if($cliente != null) {

    curl_setopt($cliente, CURLOPT_RETURNTRANSFER, 1);

    $dados = array (
      'id' => rand(1, 1000),
      'descricao' => $descricao,
      'quantidade' => intval($quantidade),
      'valor' => floatval($valor)
    );

    $novoProduto= json_encode($dados);

    curl_setopt($cliente, CURLOPT_POST, true);

    curl_setopt($cliente, CURLOPT_POSTFIELDS, $novoProduto);

    curl_exec($cliente);

    curl_close($cliente);

    header("location: index.php");

  } else {
    echo "ERRO 404";
  }

} else {
  echo "informe os valores das variaveis";
}

 ?>
