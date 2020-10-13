package main

// import "demopackage"
// import "fmt"

import (
	my_package "demopackage"
	"examples"
)

// var (
// 	a = c + b  // == 9
// 	b = f()    // == 4
// 	c = f()    // == 5
// 	d = 3      // == 5 after initialization has finished
// )

// func f() int {
// 	d++
// 	return d
// }

// func init() {
// 	fmt.Println("init", a)
// }

func main()  {
	// let's call all the functions here.
	// SomethingNotExported()
	// examples.MakeExample()
	// examples.GoToExample()
	// fmt.Println(a, b, c, d)
	// examples.SliceExample()
	examples.ContinueExample()
	my_package.HelloEveryone()
}