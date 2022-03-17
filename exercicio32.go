package main

import (
		"fmt"
)

func main() {

	for i:=65; i <=90;i++{
		fmt.Println(i)
		for j := 0; j<3; j++{
			fmt.Printf("\t%#U\n", i)
		}
	}
	
}
/* Na prática: exercício #2
Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
Por exemplo: 65 U+0041 'A' U+0041 'A' U+0041 'A' 66 U+0042 'B' U+0042 'B' U+0042 'B' ...e por aí vai.
Solução: https://play.golang.org/p/QlP6nVchq- */