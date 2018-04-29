<?php

	/* 
		É necessário que o webserver esteja rodando:
		
			http://localhost:3000/produtos

		Rest Simples feito em Golang se encontra na pasta "GoRest"

		
	*/

	echo "Cliente Rest Go App <br><br>";

	$url = "http://localhost:3000/produtos";

	echo "Carregando uma lista de produtos <br><br>";

	$cliente = curl_init($url); 

	curl_setopt($cliente, CURLOPT_RETURNTRANSFER, 1);	

	$resposta = curl_exec($cliente);

	$resultado = json_decode($resposta);
	
	$num = count($resultado);
	
	echo "$num resultados. <br><br>"; 

	for ($i = 0; $i < $num; $i++) {

		echo $resultado[$i]->id . ", ";
		echo $resultado[$i]->descricao . ", ";
		echo $resultado[$i]->valor . ", ";
		echo $resultado[$i]->quantidade . ", ";
		$status =  ($resultado[$i]->status)? "true" : "false";
		echo "$status <br><br>";

	}
?>