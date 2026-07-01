package main

import "fmt"

func naLinha(matriz [][]rune, lin int, valor rune) bool {
	for i := 0; i < len(matriz); i++ {
		if matriz[lin][i] == valor {
			return true
		}
	}
	return false
}

func naColuna(matriz [][]rune, colu int, valor rune) bool {
	for i := 0; i < len(matriz); i++ {
		if matriz[i][colu] == valor {
			return true
		}
	}

	return false
}

func quadrante(matriz [][]rune, lin, col int) []rune {
	n := len(matriz)

	var tam int

	if n == 4 {
		tam = 2
	} else {
		tam = 3
	}

	l := (lin / tam) * tam
	c := (col / tam) * tam

	q := []rune{}

	for i := 0; i < tam; i++ {
		for j := 0; j < tam; j++ {
			q = append(q, matriz[l+i][c+j])
		}
	}
	return q
}
func noQuadrante(matriz [][]rune, lin, col int, valor rune) bool {
	q := quadrante(matriz, lin, col)

	for _, x := range q {
		if x == valor {
			return true
		}
	}
	return false
}

func resolver(matriz [][]rune, index int) bool {
	n := len(matriz)

	if index == n*n {
		return true
	}

	l := index / n
	c := index % n

	if matriz[l][c] != '.' {
		return resolver(matriz, index+1)
	}

	for num := 1; num <= n; num++ {
		valor := rune('0' + num)
		if !naLinha(matriz, l, valor) && !naColuna(matriz, c, valor) && !noQuadrante(matriz, l, c, valor) {
			matriz[l][c] = valor

			if resolver(matriz, index+1) {
				return true
			}

			matriz[l][c] = '.'
		}

	}
	return false
}

func main() {
	var n int

	fmt.Scan(&n)

	matriz := make([][]rune, n)

	for i := 0; i < n; i++ {
		var linha string
		fmt.Scan(&linha)
		matriz[i] = []rune(linha)
	}

	resolver(matriz, 0)

	for i := 0; i < n; i++ {
		fmt.Println(string(matriz[i]))
	}
}
