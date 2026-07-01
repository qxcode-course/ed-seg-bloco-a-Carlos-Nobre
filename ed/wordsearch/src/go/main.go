package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(grid [][]byte, word string, i, j, pos int) bool {
	if pos == len(word) {
		return true
	}
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return false
	}

	if grid[i][j] != word[pos] {
		return false
	}

	temp := grid[i][j]
	grid[i][j] = '#'

	result := dfs(grid, word, i+1, j, pos+1) || dfs(grid, word, i-1, j, pos+1) || dfs(grid, word, i, j+1, pos+1) || dfs(grid, word, i, j-1, pos+1)

	grid[i][j] = temp

	return result

}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exists(grid [][]byte, word string) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if dfs(grid, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	if exists(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
