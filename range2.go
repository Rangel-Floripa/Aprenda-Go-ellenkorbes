package main

import "fmt"

func main() {
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		18:  "idade de ir pra festa",
	}
	fmt.Println(qualquercoisa)

	for key,value := range qualquercoisa{
		fmt.Println("quando idade for :", key, " ", value)
	}

	total := 0
	for i, _ := range qualquercoisa{
		total += i
	}
	fmt.Println(total)

	delete(qualquercoisa, 123)
	fmt.Println(qualquercoisa)

}