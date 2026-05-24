package main

import (
	"math"

	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func circulos(pen *Pen, x, y, raio float64) {

	if raio < 3 {
		return
	}

	pen.Up()
	pen.SetPosition(x, y)
	pen.Down()

	pen.DrawCircle(raio)

	qtd := 8
	novoRaio := raio * 0.34

	for i := 0; i < qtd; i++ {

		angulo := float64(i) * 2 * math.Pi / float64(qtd)

		nx := x + math.Cos(angulo)*(raio+novoRaio)
		ny := y + math.Sin(angulo)*(raio+novoRaio)

		circulos(pen, nx, ny, novoRaio)
	}
}
func arvore(pen *Pen, dist float64) {
	ang := 25.0
	fator := 0.77
	if dist < 10 {
		return
	}

	angulodir := ang - float64(randInt(-5, 5))
	anguloesq := ang + float64(randInt(-5, 5))
	pen.SetLineWidth(dist / 20)
	pen.Walk(dist)
	pen.Right(angulodir)
	arvore(pen, dist*fator)
	pen.Left(angulodir + anguloesq)
	arvore(pen, dist*fator)
	pen.Right(anguloesq)
	pen.Walk(-dist)

}
func gelo(pen *Pen, x, y, tamanho float64, nivel int) {

	if nivel <= 0 || tamanho < 4 {
		return
	}

	qtd := 5
	anguloEntre := 360.0 / float64(qtd)

	pen.Up()
	pen.SetPosition(x, y)
	pen.Down()

	for i := 0; i < qtd; i++ {

		angulo := float64(i) * anguloEntre

		pen.SetHeading(angulo)

		nx := x + math.Cos(angulo*math.Pi/180)*tamanho
		ny := y + math.Sin(angulo*math.Pi/180)*tamanho

		pen.Goto(nx, ny)

		gelo(pen, nx, ny, tamanho*0.45, nivel-1)

		pen.Up()
		pen.SetPosition(x, y)
		pen.Down()
	}
}
func quadrados(pen *Pen, x, y, tamanho float64, nivel int) {

	if nivel <= 0 || tamanho < 6 {
		return
	}

	pen.Up()
	pen.SetPosition(x, y)
	pen.Down()

	pen.DrawRect(tamanho, tamanho)

	novo := tamanho / 4
	espaco := tamanho / 10

	posicoes := [][2]float64{

		{0, 0},
		{1, 0},
		{2, 0},

		{0, 1},
		{2, 1},

		{0, 2},
		{1, 2},
		{2, 2},
	}

	cx := x + tamanho/2
	cy := y + tamanho/2

	for _, p := range posicoes {

		nx := x + p[0]*(novo+espaco)
		ny := y + p[1]*(novo+espaco)

		// desenha linha até o quadrado filho
		pen.Up()
		pen.SetPosition(cx, cy)
		pen.Down()

		pen.Goto(nx+novo/2, ny+novo/2)

		quadrados(pen, nx, ny, novo, nivel-1)
	}
}
func espiralQuadrada(pen *Pen, passos int, incremento float64) {

	tamanho := 5.0

	cores := [][3]float64{
		{255, 0, 0},
		{255, 128, 0},
		{255, 255, 0},
		{0, 255, 0},
		{0, 255, 255},
		{0, 0, 255},
		{128, 0, 255},
		{255, 0, 255},
	}

	for i := 0; i < passos; i++ {

		cor := cores[i%len(cores)]

		pen.SetRGB(cor[0], cor[1], cor[2])

		pen.Walk(tamanho)

		pen.Right(90)

		tamanho += incremento
	}
}
func triangulo(pen *Pen, x, y, tamanho float64) {

	altura := tamanho * math.Sqrt(3) / 2

	pen.Up()
	pen.SetPosition(x, y)
	pen.Down()

	pen.Goto(x+tamanho, y)

	pen.Goto(x+tamanho/2, y-altura)

	pen.Goto(x, y)
}

func sierpinski(pen *Pen, x, y, tamanho float64, nivel int) {

	if nivel <= 0 || tamanho < 4 {

		triangulo(pen, x, y, tamanho)

		return
	}

	novo := tamanho / 2

	altura := novo * math.Sqrt(3) / 2

	sierpinski(
		pen,
		x,
		y,
		novo,
		nivel-1,
	)

	sierpinski(
		pen,
		x+novo,
		y,
		novo,
		nivel-1,
	)

	sierpinski(
		pen,
		x+novo/2,
		y-altura,
		novo,
		nivel-1,
	)
}
func trigo(pen *Pen, tamanho float64, nivel int) {
	if nivel <= 0 {
		return
	}

	fatorCentro := 0.80
	fatorLado := 0.45
	angulo := 40.0

	pen.Walk(tamanho)

	pen.Left(angulo)
	trigo(pen, tamanho*fatorLado, nivel-1)

	pen.Right(angulo * 2)
	trigo(pen, tamanho*fatorLado, nivel-1)

	pen.Left(angulo)
	trigo(pen, tamanho*fatorCentro, nivel-1)

	pen.Up()
	pen.Walk(-tamanho)
	pen.Down()
}

func main() {

	pen := NewPen(1400, 1400)

	pen.SetRGB(255, 255, 255)

	pen.dc.SetRGB(0, 0, 0)
	pen.dc.Clear()

	pen.SetRGB(255, 255, 255)

	pen.SetLineWidth(1)

	circulos(pen, 700, 700, 220)

	pen.SavePNG("circulos.png")

	pen.dc.SetRGB(0, 0, 0)
	pen.dc.Clear()

	pen.SetRGB(255, 255, 255)

	pen.SetLineWidth(1)

	gelo(pen, 700, 700, 220, 5)

	pen.SavePNG("gelo.png")

	pen.dc.SetRGB(0, 0, 0)
	pen.dc.Clear()

	pen.SetRGB(255, 255, 255)

	pen.SetLineWidth(1)

	quadrados(pen, 350, 350, 700, 5)

	pen.SavePNG("quadrados.png")

	pen.dc.SetRGB(0, 0, 0)
	pen.dc.Clear()

	pen.SetRGB(255, 255, 255)

	pen.SetLineWidth(2)

	pen.Up()
	pen.SetPosition(700, 700)
	pen.SetHeading(0)
	pen.Down()

	espiralQuadrada(pen, 300, 5)

	pen.SavePNG("espiral.png")

	pen.dc.SetRGB(0, 0, 0)
	pen.dc.Clear()

	pen.SetRGB(255, 255, 255)

	pen.SetLineWidth(1)

	sierpinski(pen, 100, 1200, 1200, 7)

	pen.SavePNG("triangulos.png")

	pen.dc.SetRGB(0, 0, 0)
	pen.dc.Clear()

	pen.SetRGB(255, 255, 255)
	pen.SetLineWidth(1)

	pen.Up()
	pen.SetPosition(700, 1300)
	pen.SetHeading(90)
	pen.Down()

	trigo(pen, 200, 8)

	pen.SavePNG("trigo.png")

}
