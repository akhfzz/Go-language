package main

import (
	"fmt"
	"math"
)

type shape interface {
	luas() float64
}

type jariRuas struct {
	radius float64
}

type ukuran struct {
	lebar   float64
	panjang float64
}

func (p *ukuran) luas() float64 {
	return p.panjang * p.lebar
}

func (l *jariRuas) luas() float64 {
	return math.Phi * l.radius * l.radius
}

func getArea(g shape) float64 {
	return g.luas()
}

func main() {
	b1 := jariRuas{4.3}
	b2 := ukuran{5, 8}
	shapes := []shape{&b1, &b2}
	for _, shape := range shapes {
		//cara1//
		fmt.Println(getArea(shape))
		//cara2//
		fmt.Println(shape.luas())
	}
}
