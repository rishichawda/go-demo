package main

import "fmt"

func init() {
	fmt.Println("file scope")
}

func SomethingNotExported() {
	fmt.Println("A function in main package but in different files.")
}
