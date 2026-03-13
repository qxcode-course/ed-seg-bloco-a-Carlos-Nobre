package main

import "fmt"

func main() {
	var helicoptero, policial, ladrao, direcao int

	fmt.Scan(&helicoptero, &policial, &ladrao, &direcao)

	if direcao == -1 {
		for true {
			if ladrao == policial {
				fmt.Println("N")
				break
			} else if ladrao == helicoptero {
				fmt.Println("S")
				break
			}

			if ladrao >= 15 {
				ladrao = ladrao % 15
			}

			ladrao++
		}
	}

	if direcao == 1 {
		for true {
			if ladrao == policial {
				fmt.Println("N")
				break
			} else if ladrao == helicoptero {
				fmt.Println("S")
				break
			}

			if ladrao <= 0 {
				ladrao = 15 + ladrao
			}

			ladrao--
		}
	}
}
