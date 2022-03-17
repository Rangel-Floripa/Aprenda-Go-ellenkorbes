package main

import "fmt"

func main() {
	// Números primos de 1 a 100
	
	for i:= 0 ; i <= 100; i++{	
		if (i == 2 || i == 3 || i == 5){					
			fmt.Println(i, "é um número primo")

		} else if  i % 2 != 0 {
				if  i % 3 != 0 {
					if i % 5 != 0 {
						if i != 1 {
							fmt.Println(i, "é um número primo")
						}
					}
				}
		}
		
		
	}
}