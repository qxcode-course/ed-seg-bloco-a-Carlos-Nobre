package main

import (
	"fmt"
	"strconv"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	var sb strings.Builder

	sb.WriteString("[")

	for node := l.Front(); node != l.root; node = node.Next() {
		sb.WriteString(" ")
		sb.WriteString(strconv.Itoa(node.Value))
		if node == sword {
			sb.WriteString(">")
		}
	}
	sb.WriteString(" ]")
	return sb.String()
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if l.size == 0 {
		return nil
	}
	if it.Next() == l.root {
		return l.Front()
	}
	return it.Next()
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	//fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
