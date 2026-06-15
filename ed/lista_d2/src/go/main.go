package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	next  *Node
	prev  *Node
	root  *Node
	Value int
}

type LList struct {
	root *Node
	size int
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

func (ll *LList) String() string {
	if ll.root.next == ll.root {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	for no := ll.root.next; no != ll.root; no = no.next {
		sb.WriteString(fmt.Sprint(no.Value))

		if no.next != ll.root {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}

func (ll *LList) PushBack(value int) {
	newLast := NewNode(value)
	currentLast := ll.root.prev

	newLast.next = ll.root
	newLast.prev = currentLast
	newLast.root = ll.root
	currentLast.next = newLast
	ll.root.prev = newLast
	ll.size++

}

func (ll *LList) PushFront(value int) {
	newFist := NewNode(value)
	currentFist := ll.root.next

	newFist.prev = ll.root
	newFist.next = currentFist
	newFist.root = ll.root
	currentFist.prev = newFist
	ll.root.next = newFist
	ll.size++
}

func (ll *LList) Front() *Node {
	return ll.root.Next()
}

func (ll *LList) Back() *Node {
	return ll.root.Prev()
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) Search(value int) *Node {

	for node := ll.root.Next(); node != nil; node = node.Next() {
		if node.Value == value {
			return node
		}
	}

	return nil
}
func (ll *LList) Insert(node *Node, value int) {
	newNode := NewNode(value)
	newNode.root = ll.root
	newNode.next = node
	newNode.prev = node.prev
	newNode.prev.next = newNode

	node.prev = newNode
	ll.size++

}
func (ll *LList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	ll.size--
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
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root

	return &LList{
		root: root,
		size: 0,
	}
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func (ll *LList) Size() int {
	count := 0

	for no := ll.Front(); no != nil; no = no.Next() {
		count++
	}

	return count
}
