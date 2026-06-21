package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	data []T
}

func NewStack[T any](value int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, value),
	}
}

func (s *Stack[T]) String() string {
	output := ""
	for i := 0; i < cap(s.data); i++ {
		if i != 0 {
			output += ", "
		}
		if i < len(s.data) {
			output += fmt.Sprintf("%v", s.data[i])
		} else {
			output += "_"
		}
	}
	return output
}

func (v *Stack[T]) Push(value T) {
	v.data = append(v.data, value)
}

func (v *Stack[T]) Peek() (any, error) {
	if len(v.data) == 0 {
		return fmt.Errorf("stack is empty"), nil
	}

	last := len(v.data)
	last--
	value := v.data[last]
	return value, nil
}

func (v *Stack[T]) Size() int {
	return len(v.data)
}

func (v *Stack[T]) Clear() {
	v.data = v.data[:0]
}

func (v *Stack[T]) Pop() error {
	if len(v.data) == 0 {
		return fmt.Errorf("stack is empty")
	}

	last := len(v.data)
	last--

	v.data = v.data[:last]

	return nil
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewStack[int](10)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			cap, _ := strconv.Atoi(parts[1])
			v = NewStack[int](cap)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Push(value)
			}
		case "debug":
			fmt.Println(v.String())
		case "top":
			top, err := v.Peek()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "size":
			fmt.Println(v.Size())
		case "pop":
			err := v.Pop()
			if err != nil {
				fmt.Println(err)
			}
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
