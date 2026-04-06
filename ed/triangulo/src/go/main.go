package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processa(vet []int) {
	lista := vet;

	//? cria um vetor auxiiar que vai receber a soma dos elementos adjacentes do vetor original
	vetAux := []int{};

	//? percorre o vetor original e soma os elementos adjacentes, armazenando o resultado no vetor auxiliar
	for i := 0; i < len(lista) - 1; i++ {
		vetAux = append(vetAux, lista[i] + lista[i + 1]);
	}

	//? se o vetor auxiliar tiver pelo menos um elemento, chama a função recursivamente com o vetor auxiliar como argumento
	if len(vetAux) >= 1 {
		processa(vetAux);
	}

	//? quando o vetor auxiliar tiver apenas um elemento, imprime o resultado
	fmt.Print("[")
	fmt.Print(Join(vet, " "))
	fmt.Println("]")

	//? nao precisa retornar nada, pois a função é recursiva e o resultado final é 
	//? impresso quando o vetor auxiliar tiver apenas um elemento
	//? , ao terminar a função ele retorna automaticamente para a chamada anterior,
	//? e assim por diante, até chegar na chamada original da função no main.*/	

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	vet := []int{}
	for _, part := range parts {
		if value, err := strconv.Atoi(part); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
}

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := " "
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	s += " "
	return s
}
