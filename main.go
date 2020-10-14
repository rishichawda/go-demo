package main

// import "demopackage"
// import "fmt"

import (
	"examples"
	"fmt"
	"math/rand"
	"time"
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

// var a = b
// var b = f()
// var c = 0

// func f() int {
// 	return 4
// }

func main() {
	// let's call all the functions here.
	// SomethingNotExported()
	// examples.PointerExample()
	// examples.SliceMutationExample()
	// examples.VariadicFunction(1, 2, 3, 4, 5, 5, 6)
	// examples.MakeExample()
	// examples.GoToExample()
	// fmt.Println(a, b, c, d)
	// examples.SliceExample()
	// examples.ContinueExample()
	// my_package.HelloEveryone()

	// var mychannel = make(chan int)
	// var mychannel2 = make(chan string)
	println(examples.PanicDeferAndRecover())
	// go generate(mychannel)
	// go sayHello(mychannel2)

	// for {
	// 	select {
	// 	case value := <-mychannel:
	// 		fmt.Println(value)
	// 	case str := <-mychannel2:
	// 		fmt.Println(str)
	// 	}
	// }
	// go sayWorld()
}

func generate(ch chan int) {
	for {
		ch <- rand.Int()
		time.Sleep(time.Second / 4)
	}
	// close(ch)
}

func sayHello(ch chan string) {
	i := 0
	for {
		ch <- "hello"
		i++
		if i == 5 {
			break
		}
		time.Sleep(time.Second)
	}
}
func sayWorld() {
	for {
		fmt.Println("world")
		time.Sleep(time.Second)
	}
}
