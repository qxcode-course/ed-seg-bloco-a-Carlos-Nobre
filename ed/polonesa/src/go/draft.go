package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack struct {
	data []string
}

func (st *Stack) Push(value string) {
	st.data = append(st.data, value)
}

func (st *Stack) Pop() string {
	v := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]

	return v
}
func (st *Stack) Top() string {
	return st.data[len(st.data)-1]
}

func (st *Stack) Empty() bool {
	return len(st.data) == 0
}

func Prioridade(value string) int {
	switch value {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	}
	return 0
}

func IsOp(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "^"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	tokens := strings.Fields(scanner.Text())

	// numeros vão direto para a saida, a sequencia de operandores que vai pra saida depende da sua prioridade

	var st Stack
	var out []string

	for _, t := range tokens {
		if IsOp(t) {
			for !st.Empty() && Prioridade(st.Top()) >= Prioridade(t) {
				out = append(out, st.Pop())
			}
			st.Push(t)
		} else {
			out = append(out, t)
		}
	}

	for !st.Empty() {
		out = append(out, st.Pop())
	}
	fmt.Println(strings.Join(out, " "))
}
