package main

import "fmt"

func main() {

	x := 9
	y := 4

	switch {
	case x > 10: fmt.Println("x é maior que 10") 
	
	case x < 10, y == 4: fmt.Println("x é menor que 10 ou y é igual a quatro")
	
	case x ==10: fmt.Println("x é igual a 10")
	
	case y < 10: fmt.Println("y é menor que 10")
	}



	}

	/* Condicionais: a declaração switch
Switch:
pode avaliar uma expressão
switch statement == case (value)
default switch statement == true (bool)
não há fall-through por padrão
criando fall-through
default
cases compostos */