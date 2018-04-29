package main

/* 
	REST - Exemplo Simples Utilizando MUX

	@Author Original: github.com/gorilla/mux
	
	importar o pacote "mux" por linha de comando:

	> go get github.com/gorilla/mux

	Let's code!

	Ao executar o servidor

	Informe a url no navegador:

	localhost:3000/produtos

*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Conectando servidor...\n\n")
	
	r := mux.NewRouter().StrictSlash(true)

	// Retornar lista de Produtos
	r.HandleFunc("/produtos", getProdutos).Methods("GET")
	
	// Retorna o produto pelo ID
	r.HandleFunc("/produtos/{id}", getProduto).Methods("GET")
	
	// Cria um novo produto
	r.HandleFunc("/produtos", postProduto).Methods("POST")
	
	// Atualiza um produto pelo ID
	r.HandleFunc("/produtos/{id}", putProduto).Methods("PUT")

	// Deleta um produto pelo ID
	r.HandleFunc("/produtos/{id}", deleteProduto).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))

}

// Criando uma estrutura de Produto ~ As variaveis devem começar com a primeira letra maiuscula.
type Produto struct {
	Id 			uint64	`json:"id"`
	Descricao 	string	`json:"descricao"`
	Valor		float32	`json:"valor"`
	Quantidade	int		`json:"quantidade"`
	Status		bool	`json:"status"`	
}

// Criando um vetor de Produto
var produtos = []Produto {
	Produto{Id: 1, Descricao: "Notebook", Valor: 2700, Quantidade: 10, Status: true},
	Produto{Id: 2, Descricao: "HD 1TB", Valor: 700, Quantidade: 5, Status: true},
	Produto{Id: 3, Descricao: "Placa de Video", Valor: 1300, Quantidade: 3, Status: true},
	Produto{Id: 4, Descricao: "Fonte", Valor: 500, Quantidade: 0, Status: false},
	Produto{Id: 5, Descricao: "Processador i7", Valor: 3000, Quantidade: 0, Status: false},
}

// Função que retorna uma lista de produtos
func getProdutos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // tipo de visualização dos dados
	json.NewEncoder(w).Encode(produtos)
	fmt.Print("GET => ")
	mostrarTudo()
}

// Função que retorna um produto pelo ID
func getProduto(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	
	parametros := mux.Vars(r) // capiturando a requisição

	// Convertendo o parametro text para uint
	var base = 10 // base decimal
	var tam = 32 // tamamnho em bits

	IdProd, _ := strconv.ParseUint(parametros["id"], base, tam) // função retorna um uint64 e um erro ~ O _ (underline) ignora o erro

	// Percorrendo o vetor de produtos
	for _, produto := range produtos {

		if produto.Id == IdProd {
			json.NewEncoder(w).Encode(produto)
			fmt.Print("GET => ")
			mostrar(produto.Id)
			return
		}
	}

	fmt.Println("Produto nao encontrado...")
	fmt.Println()
	json.NewEncoder(w).Encode(&Produto{})
}

// Função para criar um novo produto
func postProduto(w http.ResponseWriter, r *http.Request) {
	
	var produto Produto

	body, _ := ioutil.ReadAll(r.Body) // lê todas as requisições que o usuario faz

	json.Unmarshal(body, &produto) // adicionando o novo elemento a estrutura

	produtos = append(produtos, produto) // adicionando o novo produto ao vetor de produtos

	/* 
		Exemplo para adicionar: 

		{
			"id": 10,
			"descricao": "Placa Mae",
			"valor": 3200.45,
			"quantidade": 15,
			"status": true
		}

	*/

	fmt.Print("POST: => ")
	fmt.Println(produto)
	fmt.Println()

}

// Função para atualizar o produto pelo Id
func putProduto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	parametros := mux.Vars(r)

	var base = 10
	var tam = 32

	IdProd, _ := strconv.ParseUint(parametros["id"], base, tam)

	for index, produto := range produtos {

		if produto.Id == IdProd {
			
			// O Produto encontrado e substituido pelo novo produto
			produtos = append(produtos[:index], produtos[index + 1:]...)

			var novoProduto Produto
			_ = json.NewDecoder(r.Body).Decode(&novoProduto)
			produtos = append(produtos, novoProduto)
			json.NewEncoder(w).Encode(novoProduto)
			fmt.Print("PUT => ")
			mostrar(novoProduto.Id)
			return 
		}
	}
	json.NewEncoder(w).Encode(produtos)
}

// Função para deletar um produto pelo Id
func deleteProduto(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	parametros := mux.Vars(r)

	var base = 10
	var tam = 32

	IdProd, _ := strconv.ParseUint(parametros["id"], base, tam)

	for index, produto := range produtos {

		if produto.Id == IdProd {
			
			fmt.Print("DELETE => ")
			mostrar(produto.Id)

			produtos = append(produtos[:index], produtos[index + 1:]...)
			
			return
		}
	}
	json.NewEncoder(w).Encode(produtos)
}

/* Funções para visualizar dados */
func mostrar(BuscaId uint64) {

	for _, produto := range produtos {
		if produto.Id == BuscaId {
			fmt.Println("Produto: {")
			fmt.Println("\t\"Id\":\t", produto.Id)
			fmt.Println("\t\"Descricao\":\t", produto.Descricao)
			fmt.Println("\t\"Valor\":\t", produto.Valor)
			fmt.Println("\t\"Quantidade\":\t", produto.Quantidade)
			fmt.Println("\t\"Status\":\t", produto.Status)
			fmt.Println("}\n")
			return
		}
	}
}

func mostrarTudo() {

	fmt.Println("Produtos: [")
	for _, produto := range produtos {
		fmt.Println("\t{")
		fmt.Println("\t\t\"Id\":\t", produto.Id)
		fmt.Println("\t\t\"descricao\":\t", produto.Descricao)
		fmt.Println("\t\t\"valor\":\t", produto.Valor)
		fmt.Println("\t\t\"quantidade\":\t", produto.Quantidade)
		fmt.Println("\t\t\"status\":\t", produto.Status)
		fmt.Println("\t}")
	}
	fmt.Println("]\n")
}