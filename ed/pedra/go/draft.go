package main
import (
    "fmt"
    "math"
    
)

func main() {
    var qntCompetidores int;

    fmt.Scan(&qntCompetidores)

    fist := make([]int, qntCompetidores)
    second := make([]int, qntCompetidores)
   

    for i := 0; i < qntCompetidores; i++{
        var primeiro,segundo int
        fmt.Scan(&primeiro , &segundo)


        
        fist[i] = primeiro
        second[i] = segundo

    }

    
    
    var campeao int 
    var maiorDiff int 

    for i := 0; i < qntCompetidores; i++{
        if(fist[i] < 10 || second[i] < 10){
            continue
        }else{
            campeao = i
            maiorDiff = int(math.Abs(float64(fist[i] - second[i])))
            break
        }
    }

    

    for i := 1; i < qntCompetidores; i++{ 
        if(fist[i] < 10 || second[i] < 10){
            continue
        }

        diff := int(math.Abs(float64(fist[i] - second[i])))

        if maiorDiff > diff{
            maiorDiff = diff
            campeao = i
        }
    }
    if campeao == 0{
        fmt.Println("sem ganhador")
    }else{
        fmt.Println(campeao)
    }


}
