<?php

 $id = addslashes(trim($_POST['id']));

if(!empty($id)) {

  $url = "http://localhost:3000/produtos/". $id;

  $cliente = curl_init($url);

  if($cliente != null) {
    curl_setopt($cliente, CURLOPT_CUSTOMREQUEST, "DELETE");

    curl_setopt($cliente, CURLOPT_POSTFIELDS, $id);

    curl_setopt($cliente, CURLOPT_RETURNTRANSFER, 1);

    curl_exec($cliente);

    curl_close($cliente);

    header("location: index.php");
  }
  else {
    echo "ERRO 404";
  }



} else {

  die("Informe o ID do produto...");

}

 ?>
