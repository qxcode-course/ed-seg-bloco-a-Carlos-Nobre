package main
import "fmt"

type Gomo struct{
    x,y int
}

func main() {
    qtdGomos := 0
    var dir rune

    fmt.Scan(&qtdGomos)
    fmt.Scanf("%c", &dir)

    cobra := make([]Gomo, qtdGomos)

    for i:= range cobra{
        fmt.Scan(&cobra[i].x, &cobra[i].y)
    }

    for i := qtdGomos - 1; i > 0; i--{
        cobra[i].x = cobra[i-1].x
        cobra[i].y = cobra[i-1].y
    }

    switch dir{
        case 'R':
            cobra[0].x++
        case 'L':
            cobra[0].x--
        case 'U':
            cobra[0].y--
        case 'D':
            cobra[0].y++
    }

    for _, gomo := range cobra{
        fmt.Println(gomo.x, gomo.y)
    }

}

