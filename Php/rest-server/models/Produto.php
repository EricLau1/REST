<?php

class Produto {

    private $con;
    private $table = "tb_produto";

    private $id;
    private $descricao;
    private $quantidade;
    private $valor;

    public function __construct($db) {
        $this->con = $db;
    }

    public function read() {

        $sql = "SELECT * FROM crud_rest.tb_produto";
        
        $stmt = $this->con->prepare($sql);

        $stmt->execute();

        return $stmt;
    
    }

    public function findOne() {
       
        $sql = "SELECT * FROM crud_rest.tb_produto WHERE id = ? ORDER BY id LIMIT 1";
        
        $stmt = $this->con->prepare($sql);
        $stmt->bindParam(1, $this->id);
        $stmt->execute();

        $row = $stmt->fetch(PDO::FETCH_ASSOC);

        $this->descricao = $row['descricao'];
        $this->quantidade = $row['quantidade'];
        $this->valor = $row['valor'];
        return $stmt;
    }

    public function create() {

        $sql = "INSERT INTO crud_rest.tb_produto (descricao, quantidade, valor) VALUES (?,?,?)";

        $stmt = $this->con->prepare($sql);

        // Limpando os dados
        $this->descricao  = htmlspecialchars(strip_tags($this->descricao));
        $this->quantidade = htmlspecialchars(strip_tags($this->quantidade));
        $this->valor      = htmlspecialchars(strip_tags($this->valor));
        

        $stmt->bindParam(1, $this->descricao);
        $stmt->bindParam(2, $this->quantidade);
        $stmt->bindParam(3, $this->valor);

        if($stmt->execute()) {
            return true;
        }

        printf("Erro: %s.\n", $stmt->error);

        return false;
    }

    public function update() {

        $sql = "UPDATE crud_rest.tb_produto SET descricao = ?, quantidade = ?, valor = ? ";
        $sql .= "WHERE id = ?";

        $stmt = $this->con->prepare($sql);

        // Limpando os dados
        $this->descricao  = htmlspecialchars(strip_tags($this->descricao));
        $this->quantidade = htmlspecialchars(strip_tags($this->quantidade));
        $this->valor      = htmlspecialchars(strip_tags($this->valor));
        $this->id         = htmlspecialchars(strip_tags($this->id));

        $stmt->bindParam(1, $this->descricao);
        $stmt->bindParam(2, $this->quantidade);
        $stmt->bindParam(3, $this->valor);
        $stmt->bindParam(4, $this->id);

        if($stmt->execute()) {
            return true;
        }

        printf("Erro: %s.\n", $stmt->error);

        return false;
    }


    public function delete() {

        $sql = "DELETE FROM crud_rest.tb_produto WHERE id = ?";

        $stmt = $this->con->prepare($sql);

        $this->id = htmlspecialchars(strip_tags($this->id));
        
        $stmt->bindParam(1, $this->id);

        if($stmt->execute()) {
            return true;
        }

        printf("Erro: %s.\n", $stmt->error);

        return false;
    }

    /* Getters and Setters */
    public function setId($id) {
        $this->id = $id;
    }
    public function getId() {
        return $this->id;
    }

    public function setDescricao($descricao) {
        $this->descricao = $descricao;
    }
    public function getDescricao() {
        return $this->descricao;
    }

    public function setQuantidade($quantidade) {
        $this->quantidade = $quantidade;
    }
    public function getQuantidade() {
        return $this->quantidade;
    }

    public function setValor($valor) {
        $this->valor = $valor;
    }
    public function getValor() {
        return $this->valor;
    }
}