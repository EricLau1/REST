<?php


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
    <h1>Novo Produto</h1>

      <form method="POST" action="createProduto.php">
      <div class="form-group">
        <label for="descricao">Descrição</label>
        <input type="text" maxlength="100" minlength="1"
        class="form-control" id="descricao" name="descricao" placeholder="Informe o nome do produto..." required>
      </div>
      <div class="form-group">
        <label for="quantidade">Qtde</label>
        <input type="number" min="1" max="1000"
        class="form-control" id="quantidade" name="quantidade" placeholder="Informe a quantidade deste produto" required>
        <small class="form-text text-muted">max. 1000</small>
      </div>
      <div class="form-group">
        <label for="valor">Valor</label>
        <input type="number" min="1" max="99999" step="any"
        class="form-control" id="valor" name="valor" placeholder="Informe o valor deste produto" required>
        <small class="form-text text-muted">Ex. 99.99</small>
      </div>

      <button type="submit" class="btn btn-primary">Salvar</button>
      <a href="index.php" class="btn btn-link">Voltar</a>
    </form>

  </div>

  <script src="bootstrap/js/bootstrap.js"></script>
</body>
</html>
