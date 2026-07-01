package main

import "fmt"

var (
	sequence []byte
	size     int
	limit    int
)

func mod(value int) int {
	if value < 0 {
		return -(value)
	}
	return value
}

func Isvalid(pos int, digit byte) bool {

	for i := pos - 1; i >= 0; i-- {
		if digit == sequence[i] && mod(pos-i) <= limit {
			return false
		}
	}

	for i := pos + 1; i < size; i++ {
		if digit == sequence[i] && mod(pos-i) <= limit {
			return false
		}
	}

	return true
}

func trasnformation(pos int) bool {
	if pos == size {
		return true
	}
	if sequence[pos] != '.' {
		if !Isvalid(pos, sequence[pos]) {
			return false
		}
		return trasnformation(pos + 1)
	}

	for digit := byte('0'); digit <= byte('0'+limit); digit++ {
		if Isvalid(pos, digit) {
			sequence[pos] = digit

			if trasnformation(pos + 1) {
				return true
			}

			sequence[pos] = '.'

		}
	}

	return false
}

func main() {
	var entry string
	fmt.Scan(&entry)
	fmt.Scan(&limit)

	sequence = []byte(entry)
	size = len(sequence)

	trasnformation(0)

	fmt.Println(string(sequence))

}
