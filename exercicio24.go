package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("%v\n%b\n%#x\n",a,a,a)
	y := a << 1
	fmt.Printf("%v\n%b\n%#x\n",y,y,y)
}

/* Na prática: exercício #4
Crie um programa que:
Atribua um valor int a uma variável
Demonstre este valor em decimal, binário e hexadecimal
Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
Demonstre esta outra variável em decimal, binário e hexadecimal */