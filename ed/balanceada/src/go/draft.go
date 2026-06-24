package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	data []rune
}

func (st *Stack) Push(value rune) {
	st.data = append(st.data, value)
}

func (st *Stack) Empty() bool {
	return len(st.data) == 0
}

func (st *Stack) Pop() rune {
	r := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return r
}

func (st *Stack) Top() rune {
	return st.data[len(st.data)-1]
}

func IsClose(value rune) bool {
	if value == ')' || value == '}' || value == ']' {
		return true
	}

	return false
}

func MyComplemt(value rune) rune {
	switch value {
	case '[':
		return ']'
	case '(':
		return ')'
	case '{':
		return '}'
	}

	return value
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	var pilha Stack

	for _, c := range s {
		if pilha.Empty() && (c == ')' || c == '}' || c == ']') {
			pilha.Push(c)
			break
		}

		if pilha.Empty() {
			pilha.Push(c)
			continue
		}

		if c == MyComplemt(pilha.Top()) {
			pilha.Pop()
		} else {
			pilha.Push(c)
		}

	}

	if pilha.Empty() {
		fmt.Println("balanceado")
	} else {
		fmt.Println("nao balanceado")
	}

}
