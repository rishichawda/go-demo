package examples

import "fmt"

type Rect struct {
	length float32
	width float32
}

type Circle struct {
	radius float32
}

type Processable interface {
	calculate()
}

func (r *Rect) calculate() {

}

func (c *Circle) calculate() {
	
}

func CalculateAreaForShape() {
	var rectangle = Rect{2, 4}
	// var rectangle2 = Rect{length: 2}
	fmt.Println(rectangle)
}