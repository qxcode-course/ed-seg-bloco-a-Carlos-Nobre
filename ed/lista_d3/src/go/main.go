package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
	l.size++
}

func equals(firstList *LList, secondList *LList) bool {
	if firstList.size != secondList.size {
		return false
	}

	for nodeOne, nodeTwo := firstList.root.next, secondList.root.next; nodeOne != firstList.root && nodeTwo != secondList.root; nodeOne, nodeTwo = nodeOne.next, nodeTwo.next {
		if nodeOne.Value != nodeTwo.Value {
			return false
		}
	}

	return true
}

func addsorted(lista *LList, value int) {
	for node := lista.root.next; node != lista.root; node = node.next {
		if node.Value > value {
			lista.insertBefore(node, value)
			return
		}
	}
	lista.PushBack(value)
}

func reverse(lista *LList) {
	if lista.size <= 1 {
		return
	}

	for node := lista.root; ; {
		node.next, node.prev = node.prev, node.next

		node = node.prev

		if node == lista.root {
			break
		}
	}
}

func merge(lista1 *LList, lista2 *LList) *LList {
	listMerged := NewLList()

	nodeA := lista1.root.next
	nodeB := lista2.root.next

	for nodeA != lista1.root && nodeB != lista2.root {
		if nodeA.Value <= nodeB.Value {
			listMerged.PushBack(nodeA.Value)
			nodeA = nodeA.next
		} else {
			listMerged.PushBack(nodeB.Value)
			nodeB = nodeB.next
		}

	}
	for nodeA != nodeA.root {
		listMerged.PushBack(nodeA.Value)
		nodeA = nodeA.next
	}

	for nodeB != nodeB.root {
		listMerged.PushBack(nodeB.Value)
		nodeB = nodeB.next
	}
	return listMerged
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

func (l *LList) String() string {
	values := []string{}

	for node := l.root.next; node != l.root; node = node.next {
		values = append(values, strconv.Itoa(node.Value))
	}
	return "[" + strings.Join(values, ", ") + "]"
}
