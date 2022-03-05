package main

import "fmt"

const (
	a0= 2022 + iota
	a1
	a2
	a3
	a4
)
func main() {


	fmt.Println(a1,a2,a3,a4)
}
/* Na prática: exercício #6
Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
Demonstre estes valores.
Solução: https://play.golang.org/p/zRBEooRvo4 */