package main

import "fmt"

func main() {
	amigos := map[string]int{
		"a": 1111111.,
		"b": 2222222.,
		"c": 3333333.,
		"d": 4444444,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["d"])

	amigos["e"]=555555

	fmt.Println(amigos)

}