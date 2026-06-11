package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type LList struct {
	root *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}
func NewLList() *LList {
	root := &Node{}

	root.next = root
	root.prev = root

	return &LList{
		root: root,
	}
}

func (ll *LList) Size() int {
	count := 0
	for n := ll.root.next; n != ll.root; n = n.next {
		count++
	}
	return count
}

func (ll *LList) PushFront(value int) {
	no := NewNode(value)

	fist := ll.root.next

	no.prev = ll.root
	no.next = fist

	fist.prev = no
	ll.root.next = no

}
func (ll *LList) PushBack(value int) {
	no := NewNode(value)

	final := ll.root.prev

	no.next = ll.root
	no.prev = final

	final.next = no
	ll.root.prev = no
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
}

func (ll *LList) PopBack() {
	if ll.root.prev == ll.root {
		return
	}

	last := ll.root.prev
	newLast := last.prev

	newLast.next = ll.root
	ll.root.prev = newLast

}

func (ll *LList) PopFront() {
	if ll.root.next == ll.root {
		return
	}

	fist := ll.root.next

	newFist := fist.next

	newFist.prev = ll.root
	ll.root.next = newFist
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

func (ll *LList) String() string {
	if ll.root.next == ll.root {
		return "[]"
	}

	var lb strings.Builder
	lb.WriteString("[")

	for n := ll.root.next; n != ll.root; n = n.next {
		lb.WriteString(fmt.Sprint(n.value))
		if n.next != ll.root {
			lb.WriteString(", ")
		}
	}
	lb.WriteString("]")

	return lb.String()
}
