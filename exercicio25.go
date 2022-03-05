package main

import "fmt"

func main() {
	a := `a  b
				c  d `
	fmt.Printf("%v\n%T", a, a)
}

/* Na prática: exercício #5
Crie uma variável de tipo string utilizando uma raw string literal.
Demonstre-a.
Solução: https://play.golang.org/p/RkpqPpRWuo  */