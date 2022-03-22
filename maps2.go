package main

import "fmt"

func main() {
	qualquercoisa := map[int]string{

		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "essa é massa",
		18:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)


	for key, value:= range qualquercoisa{
		fmt.Println(key,value)
	}

	
	delete(qualquercoisa, 123)
	fmt.Println(qualquercoisa)
}