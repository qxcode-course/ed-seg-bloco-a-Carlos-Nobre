package main

import "fmt"

func main() {
	var helicoptero, policial, ladrao, direcao int

	fmt.Scan(&helicoptero, &policial, &ladrao, &direcao)

	if direcao == -1{
		for true{
			if ladrao == helicoptero {
				fmt.Println("S")
				break
			}else if ladrao == policial{
				fmt.Println("N")
				break
			}

			ladrao--

			if ladrao < 0{
				ladrao = 15
			} 
		}

	    


	}

	if direcao == 1{
		for true{
			if ladrao == helicoptero{
				fmt.Println("S")
				break
			}else if ladrao == policial{
				fmt.Println("N")
				break
			}

			ladrao++ 

			if ladrao > 15{
				ladrao  = 0
			}
		}

	}
}
