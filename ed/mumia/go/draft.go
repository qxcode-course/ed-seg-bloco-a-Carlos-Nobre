package main

import "fmt"

func main() {

	var nome string
	fmt.Scan(&nome)

	var idade int
	fmt.Scan(&idade)

	
    if idade < 12 {

		fmt.Println(nome + " eh crianca")
		return

	} else if idade >= 12 && idade < 18 {
		fmt.Println(nome + " eh jovem")
		return

	} else if idade >= 18 && idade < 65 {
		fmt.Println(nome + " eh adulto")
		return

	} else if idade >= 65 && idade < 1000 {
		fmt.Println(nome + " eh idoso")
		return
	} else {
		fmt.Println(nome + " eh mumia")
	}

}
