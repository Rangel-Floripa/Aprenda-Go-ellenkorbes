package main

import "fmt"

func main() {
	slice := []string{"banana", "maçã", "uva"}

	for indice, valor := range slice{
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}
	
}