package main

import "fmt"

func main() {
	for hora := 0; hora <=12; hora++{
		fmt.Println("Hora", hora)
		for minuto := 0 ; minuto < 60; minuto++{
		fmt.Print(" ", minuto)
		}
		fmt.Println("")
	}
	
}