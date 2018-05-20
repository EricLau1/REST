<?php

function getConexao() {

  try {
      $pdo = new PDO("mysql:localhost;dbname=crud_rest", "root", "");
      $pdo->query("set names utf8");
      //echo "banco conectado com sucesso!";
      return $pdo;
  } catch (PDOException $e) {
    echo "Erro ao conectar => " . $e->getMessage();
  }
  return $pdo;
}
/*

$pdo = getConexao();

$sql = "SELECT * FROM crud_rest.tb_produto";

$rs = $pdo->prepare($sql);
$rs->execute();

echo "<br><br><br>";
echo "<table>";
echo "<th>Id</th> <th>Descricao</th> <th>Quantidade</th> <th>Valor</th>";

while($row = $rs->fetch(PDO::FETCH_ASSOC)) {
  echo "<tr>";
  echo "<td>". $row['id'] ."</td>";
  echo "<td>". $row['descricao'] ."</td>";
  echo "<td>". $row['quantidade'] ."</td>";
  echo "<td>". $row['valor'] ."</td>";
  echo "<tr>";
}
echo "</table>"; */
 ?>
