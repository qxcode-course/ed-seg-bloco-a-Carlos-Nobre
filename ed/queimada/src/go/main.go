package main

import (
	"bufio"
	"fmt"
	"os"
)

func burnTrees(grid [][]rune, l, c int) {
	florest, lfire, cfire := grid, l, c

	if lfire < 0 || lfire >= len(florest) || cfire < 0 || cfire >= len(florest[0]) {
		return
	}
	
	if florest[lfire][cfire] != '#' {
		return
	}

	florest[lfire][cfire] = 'o'

	burnTrees(florest, lfire+1,cfire)
	burnTrees(florest, lfire-1,cfire)
	burnTrees(florest, lfire,cfire+1)
	burnTrees(florest, lfire,cfire-1)

	// se estiver fora da matriz, retorne
	// se o elemento atual não for uma arvore, retorne
	// queime a arvore colocando o caractere 'o' na posição atual
	// chame a recursão para todos os 4 vizinhos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
