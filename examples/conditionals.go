package examples

import "fmt"

func init() {
	fmt.Println("conditionals")
}

func IfExample() {
	i := 2
	if i <= 4 {
		print(i)
	} else {
		println("i is greater than 4")
	}
}

func SwitchExample() {
	i := 4
	switch i {
	case 4:
		println("i is equal to 4")
		// fallthrough
	case 5:
		println("i is equal to 5")
	default:
		println("i is something else")
	}
}

func ifElsePattern() {
	// switch {
	// case condition1:
	// 	// something
	// case condition2:
	// 	// something else
	// default:
	// 	// final else block
	// }
}

func TypeSwitchExample() {
	// switch value := processableRect.(type) {
	// case Rect:
	// 	// If it contains value of this type
	// case Circle:
	// 	// If it contains value of this type
	// default:
	// 	// some other type
	// }
}
