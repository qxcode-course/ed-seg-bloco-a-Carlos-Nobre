package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMen(vet []int) []int {
	lista := vet
	homens := make([]int, 0)

	for _, val := range lista {
		if val >= 0 {
			homens = append(homens, val)
		}
	}
	return homens
}

func getCalmWomen(vet []int) []int {
	lista := vet
	mulheresCalmas := make([]int, 0)

	for _, val := range lista {
		if val < 0 && val >= -9 {
			mulheresCalmas = append(mulheresCalmas, val)
		}
	}
	return mulheresCalmas
}
	

func sortVet(vet []int) []int {
	lista := vet
	for i:=0; i < len(lista) - 1; i++ {
		maior := lista[i]
		for j := i + 1; j < len(lista); j++ {
			if maior > lista[j] {
				backup := lista[j]
				lista[j] = maior
				lista[i] = backup
				maior = lista[i]
			}
		}
	}
	return lista
}

func sortStress(vet []int) []int {
	lista := vet

	for i:=0; i < len(lista) - 1; i++ {
		maior := lista[i]
		for j := i + 1; j < len(lista); j++ {
			if abs(maior) > abs(lista[j]) {
				backup := lista[j]
				lista[j] = maior
				lista[i] = backup
				maior = lista[i]
			}
		}
	}

	return lista
}

func reverse(vet []int) []int {
	lista := make([]int,0)

	for i := len(vet) - 1; i >= 0; i-- {
		lista = append(lista, vet[i])
	}

	return lista
}

func unique(vet []int) []int {
	lista := make([]int, 0)
	repetidos := make(map[int]bool)
	for _, val := range vet {
		if !repetidos[val] {
			lista = append(lista, val)
			repetidos[val] = true
		}
	}
	return lista
}

func repeated(vet []int) []int {
	lista := make([]int, 0)
	repetidos := make(map[int]bool)
	for _, val := range vet {
		if repetidos[val] {
			lista = append(lista, val)
		}else{
			repetidos[val] = true
		}
	}
	return lista
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

