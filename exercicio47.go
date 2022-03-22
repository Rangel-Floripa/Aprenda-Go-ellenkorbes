package main

import "fmt"

func main() {

	ss := [][]string{
		[]string{
			"miu",
			"milton",
			"encher o saco",
		},
		[]string{
			"mimi",
			"martha",
			"pedir comida",
		},
		[]string{
			"meus alunos queridos",
			"que estudam bastante",
			"fazer os exercícios ninja",
		},
	}
	
	for _, v := range ss {
		fmt.Println(v)
		}
		
		fmt.Println("\n\n")

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}




	/* Na prática: exercício vkorbes#7
Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
"Nome", "Sobrenome", "Hobby favorito"
Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
Solução: https://play.golang.org/p/Gh81-d5tMi */
}