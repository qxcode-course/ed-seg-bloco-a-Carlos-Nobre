package main

import(
	"fmt"
	"strings"
	"strconv"
)

func mostrarjogadores(rodaMorte []bool, espada int) string {
	var builder strings.Builder

	builder.WriteString("[ ")

	for i, v := range rodaMorte {
		if !v {
			continue
		}

		builder.WriteString(strconv.Itoa(i + 1))

		if i+1 == espada {
			builder.WriteString(">")
		}

		
		builder.WriteByte(' ')

		
	}

	builder.WriteString("]")

	return builder.String()
}

func acharVivo(rodaMorte []bool, espada int) int {
	i := espada % len(rodaMorte)
	var vivo int

	for {
		if rodaMorte[i] {
			vivo = i
			break
		}
		i = (i + 1) % len(rodaMorte)
	}

	return vivo
}


func main() {
	
	var qntPessoas, espada int
	fmt.Scan(&qntPessoas, &espada)

	rodaMorte := make([]bool, qntPessoas)

	for i :=  range rodaMorte{
		rodaMorte[i] = true
	}

	for i := 0; i < qntPessoas - 1; i++ {
		fmt.Println(mostrarjogadores(rodaMorte,espada))
		vivo := acharVivo(rodaMorte, espada)
		rodaMorte[vivo] = false
		espada = (vivo + 1) % len(rodaMorte) + 1
		espada = acharVivo(rodaMorte, espada-1) + 1

		if espada > len(rodaMorte) {
			espada = 1
		}
	}

	fmt.Println(mostrarjogadores(rodaMorte,espada))

}

