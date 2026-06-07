package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func Bubble(v []int) {
	for i := 0; i < len(v); i++ {
		for j := i + 1; j < len(v); j++ {
			if Abs(v[i]) > Abs(v[j]) {
				backup := v[i]
				v[i] = v[j]
				v[j] = backup
			}
		}
	}
}

func cont(vet []int) int {

	cont := 1

	for i := 0; i < len(vet)-1; i++ {
		if Abs(vet[i]) != Abs(vet[i+1]) {
			cont++
		}
	}
	return cont
}

func occurr(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}

	Bubble(vet)

	var resp []Pair

	valor := Abs(vet[0])
	qtd := 1

	for i := 1; i < len(vet); i++ {
		if Abs(vet[i]) == valor {
			qtd++
		} else {
			resp = append(resp, Pair{valor, qtd})
			valor = Abs(vet[i])
			qtd = 1
		}
	}

	resp = append(resp, Pair{valor, qtd})

	return resp
}

func teams(vet []int) []Pair {
	var resp []Pair

	for i := 0; i < len(vet); {
		qtd := 1

		for i+qtd < len(vet) && Abs(vet[i]) == Abs(vet[i+qtd]) {
			qtd++
		}

		resp = append(resp, Pair{vet[i], qtd})
		i += qtd
	}
	return resp
}

func mnext(vet []int) []int {
	var mapa []int

	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 {
			mapa = append(mapa, 0)
			continue
		}

		temVizinhoNegativo := false

		if i > 0 && vet[i-1] < 0 {
			temVizinhoNegativo = true
		}

		if i < len(vet)-1 && vet[i+1] < 0 {
			temVizinhoNegativo = true
		}

		if temVizinhoNegativo {
			mapa = append(mapa, 1)
		} else {
			mapa = append(mapa, 0)
		}
	}
	return mapa
}

func alone(vet []int) []int {
	_ = vet
	return nil
}

func couple(vet []int) int {
	_ = vet
	return 0
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	_ = vet
	_ = seq
	return -1
}

func erase(vet []int, posList []int) []int {
	_ = vet
	_ = posList
	return nil
}

func clear(vet []int, value int) []int {
	_ = vet
	_ = value
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
