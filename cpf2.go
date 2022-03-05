package main

import (
	"fmt"
	"os"
)

var entrada [12]string
var entradaConvertida [12]int
var total1 int
var total2 int
var validador1 int
var validador2 int

func main() {

	fmt.Println("Este programa está na versão 1.2")

	fmt.Println("1- Verificar um CPF")
	fmt.Println("2- Verificar Estado de Origem")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	switch comando {
	case 1 :
		fmt.Println("Digite o número do CPF a ser verificado. Digite apenas os 11 números, separados por um espaço. Exemplo 1 1 1 1 1 1 1 1 1 1 1")
	case 2 :
		fmt.Println("Digite apenas o nono dígito do seu CPF")
	case 0 :
		fmt.Println("Saindo do programa...")
		os.Exit(1)
	default :
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
	var digito1 int64
	fmt.Sprintln(fmt.Scan(&digito1))
	
}