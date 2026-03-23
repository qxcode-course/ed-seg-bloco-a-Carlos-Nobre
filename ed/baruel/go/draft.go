package main

import (
	"fmt"
)

func main() {
    qtdAlbum := 0
    qtdFigurinhas := 0

    

    fmt.Scan(&qtdAlbum, &qtdFigurinhas)

    album := map[int]bool{}
    figurinhas := make([]int,qtdFigurinhas)
    repetidas := make([]int,0)
    faltando := make([]int,0)

    // lendo as figurinhas
    for i := range figurinhas{
        fmt.Scan(&figurinhas[i])
    }

    //criando e preencendo o album com todas as figurinhas como falsas
    for i := 1; i <= qtdAlbum; i++{
        album[i] = false
    }

    // pega a figurinhas verifica se tem no album, se tiver ele coloca como repitida, se não tiver ele vai no album e modifica de false para true
    for _, fig := range figurinhas{
        if album[fig] == false{
            album[fig] = true
        }else{
            repetidas = append(repetidas, fig)
        }
    }

    // verifica se há figurinhas repetidas, se não tiver ele printa "N", se tiver ele printa as figurinhas repetidas
    if(len(repetidas) == 0){
        fmt.Println("N")
    }else{
        for i,fig := range repetidas{
            if(i == len(repetidas)-1){
                fmt.Println(fig)
            }else{
                fmt.Printf("%d ", fig)
            }
        }
    }

    // percorre o album e verifica quais figurinhas estão faltando, se a figurinha estiver como falsa ele adiciona na lista de faltando
    for key, value := range album{
        if(value == false){
            faltando = append(faltando, key)
        }
    }

    // verifica se há figurinhas faltando, se não tiver ele printa "N", se tiver ele printa as figurinhas faltando
    if(len(faltando) == 0){
        fmt.Println("N")
    }else{
        for i,fig := range faltando{
            if(i == len(faltando)-1){
                fmt.Println(fig)
            }else{
                fmt.Printf("%d ", fig)
            }
        }
    }
}
