package examples

type Rect struct {
	length float32
	width  float32
}

type Circle struct {
	radius float32
}

type Processable interface {
	calculate()
}

func (r *Rect) calculate() {
	r.length = 3
}

func (c *Circle) calculate() {

}

func CalculateAreaForShape() {
	// var processableRect Processable
	var rectangle Rect
	rectangle.calculate()
	// processableRect = rectangle
	// // var rectangle2 = Rect{length: 2}
	// fmt.Println(processableRect.calculate())
}
