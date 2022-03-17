package main

import "fmt"

func main() {

	for i := 33; i<123; i++ {
		fmt.Printf("%d\n%#x\n%#U\n", i,i,i)
	}
	for x:=33; x<=122;x++{
		fmt.Printf("%d - %v\n", x, string(x))
	}

	/* Loops: utilizando ascii
	Desafio surpresa!
	Format printing:
	Decimal %d
	Hexadecimal %#x
	Unicode %#U
	Tab \t
	Linha nova \n
	Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.
	Solução: https://play.golang.org/p/REm2WHyzzz */
}