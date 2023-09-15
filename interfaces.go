package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Side int
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (s Square) Area() int {
	return s.Side * s.Side
}

func temp3() {

	circle := Circle{Radius: 5}
	fmt.Println("The Area of Circle is ", circle.Area())

	square := Square{Side: 5}
	fmt.Println("The Area of Square is ", square.Area())

}
