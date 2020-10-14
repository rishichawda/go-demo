package examples

import "fmt"

func MapExampleRange(names [5]string) {
	for index := range names {
		fmt.Println(names[index])
	}
}

func MakeExample() {
	var madeMap = make(map[string]int)
	madeMap["a"] = 2
	a, ok := madeMap["a"]
	println(a, ok)
}

func MapOfStructures() {
	var rectMap = make(map[string]Rect)
	rectMap = map[string]Rect{
		"1": {2, 4},
	}
	println(rectMap)
}
