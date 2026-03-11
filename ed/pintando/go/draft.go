package main

import(
    "fmt"
    "math"
)


func main() {
    
    var n1,n2,n3,p,resposta float64

    fmt.Scan( &n1, &n2,&n3)

    p = (n1 + n2 + n3) / 2

    resposta = math.Sqrt(p * (p-n1) * (p-n2) * (p-n3))

    fmt.Printf("%.2f\n", resposta)


}
