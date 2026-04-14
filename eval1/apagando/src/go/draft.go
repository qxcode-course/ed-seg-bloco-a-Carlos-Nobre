package main
import "fmt"
func main() {
	var tamFila,qntDesertores = 0,0 

	fmt.Scanf("%d", &tamFila)

	fila := make([]int, tamFila)

	for i := 0; i < tamFila; i++ {
		fmt.Scanf("%d", &fila[i])
	}

	fmt.Scanf("%d", &qntDesertores);
	desertores := make([]int, qntDesertores);

	for i := 0; i < qntDesertores; i++ {
		fmt.Scanf("%d", &desertores[i])
	}

	filaFinal := make(map[int]bool)

	for _, cara := range fila {
		filaFinal[cara] = true
	}

	for _, desertor := range desertores {
		filaFinal[desertor] = false
	}

	for _,cara := range fila{
		if filaFinal[cara] {
			fmt.Printf("%d ", cara)
		}
	}
	

}
