package main

import "fmt"

func main(){
	numerobytes, err := fmt.Println("Olá Mundo!", 100)
	fmt.Println(numerobytes, err)
}