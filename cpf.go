package main

import (
	"fmt"
	"os"
	"strconv"
)
		var cpfdigitado[12]string
		var cpfConvertido[12]int
		var total1 int 
		var total2 int
		var verificador1 int
		var verificador2 int

func main() {
	if len(os.Args) < 3 {
	fmt.Println("Para vefificar a autenticidade de um número de CPF, digite os 11 números que o compõe, separados por um espaço, após o comando de execução do arquivo.  Exemplo: go run cpf.go 1 1 1 1 1 1 1 1 1 1 1")
	os.Exit(1)
	}

	for i := 1; i < 12; i++ {
		cpfdigitado[i] = os.Args[i]
	}
			
	for i, v := range cpfdigitado {
		digitoConvertido, _ := strconv.Atoi(v)
		cpfConvertido[i] = digitoConvertido
	}
	
	if cpfConvertido[1] == cpfConvertido[2] && cpfConvertido[2] == cpfConvertido[3] && cpfConvertido[3] == cpfConvertido[4] && cpfConvertido[4] == cpfConvertido[5] && cpfConvertido[5] == cpfConvertido[6] && cpfConvertido[6] == cpfConvertido[7] && cpfConvertido[7] == cpfConvertido[8] && cpfConvertido[8]== cpfConvertido[9] && cpfConvertido[9] == cpfConvertido[10] && cpfConvertido[10] == cpfConvertido[11] {
		fmt.Println("Não se trata de um CPF válido!!! A RF não aceita a sequência de dígitos iguais como CPF válido, embora o algoritmo criado o valide")
		os.Exit(1)
	}
	

		
	b := 10
	for a :=1; a < 10 ;a++ {
		total1 += cpfConvertido[a]*b
		b -= 1
	}	
		var v1 int
		v1 = (total1*10) % 11
		
		if v1 > 9 {
			verificador1 = 1
		} else { 
			verificador1 = v1
		}

	c := 11
	for d :=1; d < 11 ;d++ {
		total2 += cpfConvertido[d]*c
		c -= 1
	}
		var v2 int
		v2 = (total2*10) % 11
		if v2 > 9 {
			verificador2 = 1
		} else {
			verificador2 =v2
		}



	if verificador1 == cpfConvertido[10] {

	} else {
		fmt.Println("O número informado não é um cpf válido. Divergência no primeiro dígito verificador. Digitado:", cpfConvertido[10], "real:", verificador1)
		os.Exit(1)
	}
	if verificador2 == cpfConvertido[11] {

	} else {
		fmt.Println("O número informado não é um cpf válido! Divergência no segundo dígito verificador. Digitado:", cpfConvertido[11], "real:", verificador2)
		os.Exit(1)

	}
	
	fmt.Println("CPF Valido!!!!")

	
	fmt.Print("O nono dígito do CPF indica a origem, neste caso, gerado no(a)  ")
	switch cpfConvertido[9]{
	case 0: fmt.Println("RS")
	case 1: fmt.Println("DF , GO, MA, MS ou TO")
	case 2: fmt.Println("AM, PA, RR, AP, AC ou RO")
	case 3: fmt.Println("CE, MA ou PI")
	case 4: fmt.Println("PB, PE, AL ou RN")
	case 5: fmt.Println("BA ou SE")
	case 6: fmt.Println("MG")
	case 7: fmt.Println("RJ ou ES")
	case 8: fmt.Println("SP")
	case 9: fmt.Println("PR ou SC")
	}

	

}
	


