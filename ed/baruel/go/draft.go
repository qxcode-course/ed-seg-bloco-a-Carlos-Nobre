package main
import(
    "fmt"
)
func main() {
    var tamAlbum,qntFiguras int

    fmt.Scan(&tamAlbum, &qntFiguras)

    figurinhas := make([]int, qntFiguras)
    album := make([]int, tamAlbum)

    for i:= 0; i < qntFiguras; i++{
        fmt.Scan(&figurinhas[i])
    }

    
}
