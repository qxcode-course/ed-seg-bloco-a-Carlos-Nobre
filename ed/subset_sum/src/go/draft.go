package main

import "fmt"

func subset_sum(target int, index int, elements []int, sum int) bool {
	if sum == target {
		return true
	}

	if index == len(elements)-1 {
		return false
	}

	index++

	return subset_sum(target, index, elements, sum+elements[index]) || subset_sum(target, index, elements, sum)

}

func main() {

	var n, target int
	fmt.Scan(&n, &target)

	elements := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&elements[i])
	}

	if subset_sum(target, -1, elements, 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
