package main
import "fmt"

func main() {
	var numeroPessoas,numeroSaidas int
	fmt.Scan(&numeroPessoas)

	fila := make([]int, numeroPessoas)

	for i := range fila{
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&numeroSaidas)

	pessoasSairam := make(map[int]bool, numeroSaidas)

	for i := 0; i < numeroSaidas; i++ {
		var pessoa int
		fmt.Scan(&pessoa)
		pessoasSairam[pessoa] = true
	}


	for _, valor := range fila {
		if pessoasSairam[valor]{
			continue
		}else{
			fmt.Print(valor)
			fmt.Print(" ")
		}
	}

	fmt.Println("")

}
