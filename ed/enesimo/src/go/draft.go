package main

import "fmt"

func IsPrimo(x int, div int) bool {
	_, _ = x, div
	if x < 2 {
		return false
	}
	if div == x {
		return true
	}
	if x%div == 0 {
		return false
	}

	return IsPrimo(x, div+1)

}
func Nesimo(value int, test int, count int) int {

	if count == value {
		return test - 1
	}
	if IsPrimo(test, 2) {
		count++
	}

	return Nesimo(value, test+1, count)
}
func main() {
	var x int
	fmt.Scan(&x)

	result := Nesimo(x, 1, 0)

	fmt.Println(result)
}
