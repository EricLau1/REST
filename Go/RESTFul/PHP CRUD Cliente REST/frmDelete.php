<?php

$id = addslashes(trim($_GET['id']));

if(!empty($id)) {

  $url = "http://localhost:3000/produtos/". $id;

  $client = curl_init($url);

  curl_setopt($client, CURLOPT_RETURNTRANSFER, 1);

  $response = curl_exec($client);

  //echo $response;

  $rs = json_decode($response);

  if(!$rs) {
    die('Erro 404');
  }

} else {
  header("location: index.php");
  die("Erro!");
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
    <h1>Deletar Produto</h1>
    <div class="row">
      <div class="col-md-6">
        <form method="POST" action="deleteProduto.php">

          <div class="form-group">
            <label for="id">Id</label>
            <input type="number" max="999999" min="1" value="<?php echo $id; ?>"
            class="form-control" id="id" name="id" readonly>
          </div>
          <div class="form-group">
            <label for="descricao">Descrição</label>
            <input type="text" maxlength="100" minlength="1" value="<?php echo $rs->descricao; ?>"
            class="form-control" id="descricao" name="descricao"  readonly>
          </div>
          <div class="form-group">
            <label for="quantidade">Qtde</label>
            <input type="number" min="1" max="1000" value="<?php echo $rs->quantidade; ?>"
            class="form-control" id="quantidade" name="quantidade"  readonly>
            <small class="form-text text-muted">max. 1000</small>
          </div>
          <div class="form-group">
            <label for="valor">Valor</label>
            <input type="number" min="1" max="99999" step="any" value="<?php echo $rs->valor; ?>"
            class="form-control" id="valor" name="valor"  readonly>
            <small class="form-text text-muted">Ex. 99.99</small>
          </div>

          <button type="submit" class="btn btn-danger">Deletar</button>
          <a href="index.php" class="btn btn-link">Voltar</a>
        </form>
      </div>
    </div>

  </div>

  <script src="bootstrap/js/bootstrap.js"></script>
</body>
</html>
