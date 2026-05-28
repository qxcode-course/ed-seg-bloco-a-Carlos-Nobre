package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	size     int
	capacity int
	data     []int
}

func (ms *MultiSet) Order() {
	for i := 0; i < ms.size-1; i++ {
		for j := i + 1; j < ms.size; j++ {
			if ms.data[i] > ms.data[j] {
				backup := ms.data[i]
				ms.data[i] = ms.data[j]
				ms.data[j] = backup
			}
		}
	}
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		capacity: capacity,
		size:     0,
		data:     make([]int, capacity),
	}
}

func (ms *MultiSet) Show() string {
	var str strings.Builder

	str.WriteRune('[')

	for i := 0; i < ms.size; i++ {
		str.WriteString(strconv.Itoa(ms.data[i]))

		if i == ms.size-1 {
			continue
		}

		str.WriteString(", ")
	}

	str.WriteRune(']')

	return str.String()
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (ms *MultiSet) Insert(value int) {
	if ms.capacity == 0 {
		ms.capacity = 1
	}
	if ms.size >= ms.capacity {
		ms.capacity *= 2
	}

	newMs := make([]int, ms.capacity)

	for i := 0; i < ms.size; i++ {
		newMs[i] = ms.data[i]
	}

	newMs[ms.size] = value
	ms.data = newMs
	ms.size++

	ms.Order()

}

func (ms *MultiSet) Contains(value int) bool {
	small := 0
	tall := ms.size - 1
	try := 0

	for small <= tall {
		mid := (small + tall) / 2
		try = ms.data[mid]

		if try == value {
			return true
		}
		if try > value {
			tall = mid - 1
		} else {
			small = mid + 1
		}
	}

	return false
}

func (ms *MultiSet) Erase(value int) error {
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			ms.data[i] = 0

			for j := i; j < ms.size-1; j++ {
				ms.data[j] = ms.data[j+1]

			}
			ms.size--
			ms.data[ms.size] = 0
			return nil
		}
	}

	return fmt.Errorf("value not found")
}

func (ms *MultiSet) count(value int) int {
	count := 0

	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			count++
		}
	}

	return count
}

func (ms *MultiSet) Unique() int {
	count := 0
	comp := 0
	for i := 0; i < ms.size; i++ {
		if ms.data[i] != comp {
			count++
			comp = ms.data[i]
		}
	}

	return count
}
func (ms *MultiSet) Clear() {
	for i := 0; i < ms.size; i++ {
		ms.data[i] = 0
	}
	ms.size = 0
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.Show())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			err := ms.Contains(value)
			if err {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			err := ms.count(value)
			fmt.Println(err)
		case "unique":
			err := ms.Unique()
			fmt.Println(err)
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
