package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	//fmt.Printf("%v %v %v", x, y, z)
	s := fmt.Sprintf("%v \t%v \t%v",x, y, z)
	fmt.Println(s)

}

/* Utilizando a solução do exercício anterior:
Em package-level scope, atribua os seguintes valores às variáveis:
para "x" atribua 42
para "y" atribua "James Bond"
para "z" atribua true
Na função main:
Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
Demonstre a variável "s".
Solução: https://play.golang.org/p/QFctSQB_h3 */