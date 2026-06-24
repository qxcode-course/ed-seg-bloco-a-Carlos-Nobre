package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

var dl = []int{-1, 1, 0, 0}
var dc = []int{0, 0, -1, 1}

func runnig(grid [][]rune)

func main() {
	var lines, colums int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &lines, &colums)

	labirinto := make([][]rune, lines)

	for i := 0; i < lines; i++ {
		scanner.Scan()
		labirinto[i] = []rune(scanner.Text())
	}
	showLabirinto(labirinto)
}

func showLabirinto(labirinto [][]rune) {
	for _, linha := range labirinto {
		fmt.Println(string(linha))
	}
}
