package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	vet, valor := slice, value
	baixo := 0
	alto := len(vet) - 1
	chute := 0
	meio := 0

	for baixo <= alto {
		meio = (baixo + alto) / 2
		chute = vet[meio]

		if chute == valor {
			return true, meio
		}
		if chute > valor {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	if chute > valor {
		return false, meio
	}

	return false, chute

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
