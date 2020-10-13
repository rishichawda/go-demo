package examples

import "fmt"

func forLoopExample(names []string) {
	for index := 0; index < 5; index++ {
		fmt.Println(names[index])
	}
}

func whileLoopEquivalent() {
	index := 0
	for index < 5 {
		// something
		index += 1
	}
	println(index)
}

func infiniteLoop() {
	index := 0
	for {
		// something
	}
	println(index)
}