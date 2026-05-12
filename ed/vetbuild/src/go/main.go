package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (v *Vector) PushBack(value int) {
	if v.capacity == v.size {
		newCapacity := v.capacity * 2
		newData := make([]int, newCapacity)

		for i, valor := range v.data {
			newData[i] = valor
		}

		newData[v.size] = value
		v.size++
		v.capacity = newCapacity
		v.data = newData
	} else {
		v.data[v.size] = value
		v.size++
	}

}

func (v *Vector) Status() string {
	var b strings.Builder

	b.WriteString("size:")
	b.WriteString(strconv.Itoa(v.size))
	b.WriteRune(' ')
	b.WriteString("capacity:")
	b.WriteString(strconv.Itoa(v.capacity))

	return b.String()

}

func (v *Vector) Show() string {
	var s strings.Builder

	s.WriteRune('[')

	for i := 0; i < v.size; i++ {
		if i == v.size-1 {
			s.WriteString(strconv.Itoa(v.data[i]))
		} else {
			s.WriteString(strconv.Itoa(v.data[i]))
			s.WriteRune(',')
			s.WriteRune(' ')
		}
	}

	s.WriteRune(']')

	return s.String()
}

func (v *Vector) At(index int) (string, error) {
	if index <= v.size-1 {
		return strconv.Itoa(v.data[index]), nil
	}
	return "", fmt.Errorf("index out of range")
}

func (v *Vector) Set(index int, value int) error {
	if index > v.size-1 {
		return fmt.Errorf("index out of range")
	} else {
		v.data[index] = value
	}

	return nil
}

func (v *Vector) Clear() {
	for i := 0; i < v.size; i++ {
		v.data[i] = 0
	}
	v.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
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
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v.Show())
		case "status":
			fmt.Println(v.Status())
		case "pop":
			// err := v.PopBack()
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "insert":
			// index, _ := strconv.Atoi(parts[1])
			// value, _ := strconv.Atoi(parts[2])
			// err := v.Insert(index, value)
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "erase":
			// index, _ := strconv.Atoi(parts[1])
			// err := v.Erase(index)
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "indexOf":
			// value, _ := strconv.Atoi(parts[1])
			// index := v.IndexOf(value)
			// fmt.Println(index)
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
			// if v.Contains(value) {
			// 	fmt.Println("true")
			// } else {
			// 	fmt.Println("false")
			// }
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			// newCapacity, _ := strconv.Atoi(parts[1])
			// v.Reserve(newCapacity)
		case "slice":
			// start, _ := strconv.Atoi(parts[1])
			// end, _ := strconv.Atoi(parts[2])
			// slice := v.Slice(start, end)
			// fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
