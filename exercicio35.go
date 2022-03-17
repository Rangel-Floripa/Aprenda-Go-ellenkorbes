package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println("Número",i,";resto da divisão por 4 =",i % 4)
	}

}

/* Na prática: exercício vkorbes#5
Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
Solução: https://play.golang.org/p/zcEsXqnBr8 */