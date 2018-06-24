<?php

class Database {

    private $con;

    public function conectar() {

        $this->con = null;

        try {

            $this->con = new PDO("mysql:host=localhost;dbname=crud_rest", "root", "");
            $this->con->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        } catch (PDOException $e) {

            echo "Erro ao conectar : " . $e->getMessage();

        }
        return $this->con;
    }
}
