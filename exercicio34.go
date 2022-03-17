package main

import "fmt"

func main() {
	anoQueNasci := 1965
	anoQueQueroContar := 2021

	for {
		if anoQueNasci >= anoQueQueroContar {
			break	
		}
		fmt.Println(anoQueNasci)
			anoQueNasci++
	}

}

/* Na prática: exercício vkorbes#4
Crie um loop utilizando a sintaxe: for {}
Utilize-o para demonstrar os anos desde que você nasceu.
Solução: https://play.golang.org/p/dIbfdiuw0ms */