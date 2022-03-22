package main

import "fmt"

func main() {

	array := [5]int{0, 1, 2, 3, 4}

	for i , v := range array {
		fmt.Println(i , v)
	}
	fmt.Printf("%T", array)

/* 09 – Exercícios: Ninja Nível 4
Na prática: exercício vkorbes#1
Usando uma literal composta:
Crie um array que suporte 5 valores to tipo int
Atribua valores aos seus índices
Utilize range e demonstre os valores do array.
Utilizando format printing, demonstre o tipo do array.
Solução: https://play.golang.org/p/tpWDzzlO2l 
*/

}