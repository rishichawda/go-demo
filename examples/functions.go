package examples

import "fmt"

func init() {
	fmt.Println("function")
}

func NormalFunctionWithParameters(x int, y int) int {
	sum := x + y
	return sum
}

func NamedReturn(x int, y int) (sum int) {
	sum = x + y
	return
}

func VariadicFunction(numbers ...int) {
	for index, n := range numbers {
		println(index, n)
	}
}

func fibonacciClosure() func() int {
	a, b, c := 0, -1, -1
	return func() int {
		if c < 0 {
			b += 1
			c = c + b
			return b
		} else {
			c = a + b
			a, b = b, c
			return c
		}
	}
}

func Fibonacci() {
	f := fibonacciClosure()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
