<?php

 $id = addslashes(trim($_POST['id']));
 $desc = addslashes(trim($_POST['descricao']));
 $qtd = addslashes(trim($_POST['quantidade']));
 $val = addslashes(trim($_POST['valor']));

//echo "$id, $desc, $qtd, $val";

if(!empty($id) && !empty($desc) && !empty($qtd) && !empty($val)) {

  $url = "http://localhost:3000/produtos/". $id;

  $cliente = curl_init($url);

  if($cliente != null) {

    $dados = array (
      'id' => intval($id),
      'descricao' => $desc,
      'quantidade' => intval($qtd),
      'valor' => floatval($val)
    );

    $novoProduto= json_encode($dados);

    curl_setopt($cliente, CURLOPT_CUSTOMREQUEST, "PUT");

    curl_setopt($cliente, CURLOPT_POSTFIELDS, $novoProduto);

    curl_setopt($cliente, CURLOPT_RETURNTRANSFER, 1);

    curl_exec($cliente);

    curl_close($cliente);

    header("location: index.php");
  }
  else {
    echo "ERRO 404";
  }

} else {

  echo "Informe os valores das variáveis";

}

 ?>
