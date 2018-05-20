<?php

    $url = "http://localhost:3000/produtos";

    $client = curl_init($url);

    curl_setopt($client, CURLOPT_RETURNTRANSFER, 1);

    $response = curl_exec($client);

    //echo $response;

    $rs = json_decode($response);

    if(!$rs) {
      die('Erro 404');
    }

 ?>
<!DOCTYPE html>
<html lang="pt-br">
<head>
  <meta charset="utf-8"/>
  <title>Produto Rest</title>
  <link rel="stylesheet" type="text/css" href="bootstrap/css/bootstrap.css" />

  <link href="https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">

</head>
<body>

  <div class="container">
    <h1>Produtos</h1>
    <div class="row">
      <div class="col-md-12">
        <a href="frmCreate.php" class="btn btn-link">Novo produto</a>
      </div>
    </div>

    <div class="row">
      <div class="col-md-12">
        <?php

            echo "<table class='table'>";
            echo "<th> ID </th> <th> Descrição </th> <th> Qtde. </th> <th> Valor </th>";
            echo "<th> Editar </th> <th> Excluir </th>";

            for($i = 0; $i < count($rs); $i++) {
              echo "<tr>";
              
              $id = $rs[$i]->id;
              echo "<td>" . $id . "</td>";
              echo "<td>" . $rs[$i]->descricao . "</td>";
              echo "<td>" . $rs[$i]->quantidade . "</td>";
              echo "<td>" . $rs[$i]->valor . "</td>";

              echo "<td><a href='frmUpdate.php?id=$id' class='btn btn-outline-primary'> <i class='fa fa-pencil'></i> </a></td>";
              echo "<td><a href='frmDelete.php?id=$id' class='btn btn-outline-danger'><i class='fa fa-trash-o'></i></a></td>";

              echo "<tr>";

            }
            echo "</table>";


         ?>

      </div>
    </div>
    <?php
    ?>
  </div>

  <script src="bootstrap/js/bootstrap.js"></script>
</body>
</html>
