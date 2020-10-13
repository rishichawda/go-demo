package examples

func PointerExample()  {
	var str = "hello"
	var pointer *string = &str
	// var rectangleObject = Rect{2,4}
	// var pointerRect = &rectangleObject
	println(*pointer)
	// println(*pointer, pointerRect.length)
}