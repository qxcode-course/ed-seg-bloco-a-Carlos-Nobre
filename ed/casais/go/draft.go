package main

import (
	"fmt"
)
func main() {
    var qntAnimais int

    fmt.Scan(&qntAnimais)

    animais := make([]int, qntAnimais)
    
    
    for i:=0; i < qntAnimais; i++{
        fmt.Scan(&animais[i])
    }

    var pares int;

    for i:=0; i < qntAnimais; i++{
       for j := i+1 ; j < qntAnimais; j++{
            if(animais[i] == 0){
                continue
            }
            
            if animais[i] + animais[j] == 0{
                animais[i] = 0
                animais[j] = 0
                pares++
            }
       }
    
    }

    fmt.Println(pares)
}
