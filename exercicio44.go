package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	y := []int{56,57,58,59,60}
	x = append(x, y...)
	fmt.Println(x)


	/* Na prática: exercício vkorbes#4
	Começando com a seguinte slice:
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	Anexe a ela o valor 52;
	Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
	Demonstre a slice;
	Anexe a ela a seguinte slice:
	y := []int{56, 57, 58, 59, 60}
	Demonstre a slice x.
	Solução: https://play.golang.org/p/6WNJ0Otpy0 */
}