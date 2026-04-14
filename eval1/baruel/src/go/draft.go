package main
import "fmt"
func main() {
    var tamAlbum, qntFigurinhas int;
    fmt.Scan(&tamAlbum, &qntFigurinhas);
    
    album := make(map[int]bool);
    minhasFigurinhas := make([]int, qntFigurinhas);
    repetidas := make([]int, 0);

    //preenche as figurinhas do user    
    for i := 0; i < qntFigurinhas; i++ {
        fmt.Scan(&minhasFigurinhas[i])
    }
    //preenche os indices do album com false
    for i := 0; i < tamAlbum; i++ {
        album[i + 1] = false;
    }


    //verifica se no album tem a figurinha, se tiver false vira true, se tiver true vira repetida
    for _,v := range minhasFigurinhas{
        if album[v] == false {
            album[v] = true;
        } else {
            repetidas = append(repetidas, v);
        }
    }

    //imprime as figurinhas repetidas ou N se não tiver nenhuma
    if len(repetidas) == 0{
        fmt.Println("N")
    } else {
        for i,v := range repetidas{
            if i == len(repetidas) - 1 {
                fmt.Println(v)
                continue;
            }
            fmt.Printf("%d ", v)
            
        }

    }

    //imprime as figurinhas que faltam ou N se não tiver nenhuma
    
    faltando := make([]int, 0);
    for i,v := range album {
        if v == false{
            faltando = append(faltando, i)
        }
    }
    //ordena as figurinhas que faltam
    for i := 0; i < len(faltando); i++ {
        for j := i + 1; j < len(faltando); j++ {
            if faltando[i] > faltando[j] {
                faltando[i], faltando[j] = faltando[j], faltando[i]
            }
        }
    }

    if len(faltando) == 0{
        fmt.Println("N")
    } else {
        for i,v := range faltando{
            if i == len(faltando) - 1 {
                fmt.Println(v)
                continue;
            }
            fmt.Printf("%d ", v)
        }
    }
}
