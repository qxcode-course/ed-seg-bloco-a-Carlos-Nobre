package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func arvore(pen *Pen, dist float64) {
	ang := 25.0
	fator := 0.77
	if dist < 10 {
		return
	}

	
	angulodir := ang - float64(randInt(-5,5))
	anguloesq := ang + float64(randInt(-5,5))
	pen.SetLineWidth(dist / 20)
	pen.Walk(dist)

	if dist > 10 && dist < 50 {
		pen.DrawCircle(10.0)
		pen.SetRGB(48,0,0)
	}
	
	pen.Right(angulodir)
	arvore(pen, dist * fator)
	pen.Left(angulodir + anguloesq)
	arvore(pen, dist * fator)
	pen.Right(anguloesq)
	pen.Walk(-dist)
	pen.SetRGB(3, 70,17)

}

func main() {
	pen := NewPen(1000, 1000)   // cria um canvas de 1000 de largura por 1000 de altura
	pen.SetRGB(3, 70,17)  // muda a cor do pincel para vermelho
	pen.SetPosition(500, 1000) // move o pincel para x 250, y 500
	pen.SetHeading(90)

	arvore(pen, 200.0)
	
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
