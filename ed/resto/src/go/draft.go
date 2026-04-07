package main
import "fmt"

func divisao(valor int) {
    if valor == 0 {
        return
    }

    resto := valor % 2
    valor = valor / 2
    divisao(valor)

    fmt.Println(valor, resto);


}

func main() {
    var valor int
    fmt.Scan(&valor)
    divisao(valor)
}
