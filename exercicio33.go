package main

import "fmt"

func main() {
	/* for i := 1965; i <= 2021; i++ {
		fmt.Println(i)
	} */
	anoQueNasci := 1965
	anoQueQueroContar := 2021

	for anoQueNasci <= anoQueQueroContar {
		fmt.Println(anoQueNasci)
		anoQueNasci++
	}


}

/* Na prática: exercício vkorbes#3
Crie um loop utilizando a sintaxe: for condition {}
Utilize-o para demonstrar os anos desde que você nasceu.
Solução: https://play.golang.org/p/qnFjiDJzLor */