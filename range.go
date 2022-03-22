package main

import "fmt"

func main() {
	slice := []string{"banana", "maçã", "uva","pessego"}
	slice[3]="melancia"
	slice = append(slice,"manga")


	for indice, valor := range slice{
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}


	slice2 :=  []int{1,2,3,4,5,6,7,8,9}
	total := 0
	for _ , valor := range slice2{
		total += valor
	}
	fmt.Println(total)
	
}