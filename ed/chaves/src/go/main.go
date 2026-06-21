package main

import "fmt"

func main() {
	fila := NewQueue[string]()

	for letra := 'A'; letra <= 'P'; letra++ {
		fila.Enqueue((string(letra)))
	}

	for i := 0; i < 15; i++ {
		var golsA, golsB int

		fmt.Scan(&golsA, &golsB)

		timeA := fila.Dequeue()
		timeB := fila.Dequeue()

		if golsA > golsB {
			fila.Enqueue(timeA)
		} else {
			fila.Enqueue(timeB)
		}
	}

	fmt.Println(fila.items.Front().Value)

}
