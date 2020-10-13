package main

import "fmt"

func SomethingNotExported()  {
	fmt.Println("A function in main package but in different files.")
}