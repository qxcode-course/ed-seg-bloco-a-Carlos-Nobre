package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minRec(vet []int, i int) int {
	if i == 0{
		return 0
	}

	menor := minRec(vet, i - 1)

	if vet[i] < vet[menor]{
		return i
	}
	return menor
}

func multRec(vet []int, i int) int {
		if i == 0{
			return vet[0]
		} 
		
		return vet[i] * multRec(vet, i - 1)
}

func tostr(vet []int) string {
	lista := vet
	var s strings.Builder
	s.WriteByte('[')

	for i,v := range lista {
		if i > 0{
			s.WriteString(", ")
		}
		s.WriteString(strconv.Itoa(v))
	}
	s.WriteByte(']')


	return s.String()
}

func tostrrev(vet []int) string {
	lista := vet
	var s strings.Builder
	s.WriteByte('[')

	for i := len(lista) - 1; i >= 0; i-- {
		if i  < len(lista) - 1{
			s.WriteString(", ")
		}
		s.WriteString(strconv.Itoa(lista[i]))
	}
	s.WriteByte(']')


	return s.String()
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	lista := vet
	contador := 1

	for i := 0; i < len(lista)/2; i++{
		var backup int
		backup = lista[i]
		lista[i] = lista[len(lista) - i - 1]
		lista[len(lista) - i - 1] = backup
		contador++
	}

}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	var soma int

	for _, v := range vet {
		soma += v
	}

	return soma
}


// mult: produto dos elementos do slice
func mult(vet []int) int {
	if(len(vet) == 0){
		return 1
	}
	return multRec(vet, len(vet) - 1)
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}
	return minRec(vet, len(vet) - 1)
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
