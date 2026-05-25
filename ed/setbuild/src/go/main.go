package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type Set struct {
	head     *Node
	size     int
	capacity int
}

func newSet(capacity int) *Set {
	return &Set{
		head:     nil,
		size:     0,
		capacity: capacity,
	}
}

func (s *Set) show() string {
	var str strings.Builder
	str.WriteRune('[')

	atual := s.head

	for atual != nil {
		str.WriteString(strconv.Itoa(atual.data))

		if atual.next != nil {
			str.WriteString(", ")
		}

		atual = atual.next
	}

	str.WriteRune(']')

	return str.String()
}

func (s *Set) contains(value int) bool {
	current := s.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}
func (s *Set) Insert(value int) {
	//verifica se o valor esta presente na lista
	if s.contains(value) {
		//fmt.Println("fail: valor repetido")
		return
	}
	//aumenta o tamanho da lista caso necessário
	if s.size >= s.capacity {
		if s.capacity == 0 {
			s.capacity = 1
		} else {
			s.capacity *= 2
		}
	}

	new := &Node{
		data: value,
		next: nil,
	}

	if s.head == nil || value < s.head.data {
		new.next = s.head
		s.head = new
		s.size++
		return
	}

	current := s.head

	for current.next != nil && current.next.data < value {
		current = current.next
	}

	new.next = current.next
	current.next = new

	s.size++
}

func (s *Set) erase(value int) error {
	if s.head == nil {
		return fmt.Errorf("value not found")
	}

	if s.head.data == value {
		s.head = s.head.next
		s.size--
		return nil
	}

	current := s.head.next
	before := s.head

	for current.next != nil {
		if current.data == value {
			before.next = current.next
			current.next = nil
			return nil
		} else {
			current = current.next
			before = before.next
		}
	}

	return fmt.Errorf("value not found")
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := newSet(0)
	for scanner.Scan() {
		fmt.Print("$")
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
			value, _ := strconv.Atoi(parts[1])
			v = newSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.show())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
