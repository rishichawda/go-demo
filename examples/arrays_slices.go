package examples

import "fmt"

func init() {
	fmt.Println("array")
}

func ArrayExample() {
	// var arr [5]int
	var arr = make([]int, 5)
	arr[0] = 2
}

func SliceExample() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var slice []int = arr[1:3]
	slice = append(slice, 2)
	slice2 := append(slice, 9)
	// var slice = make([]int, 0, 10)
	println(slice[len(slice)-1], slice2[len(slice2)-1])
	// println(slice)
}

func SliceMutationExample() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var slice []int = arr[1:3]
	var slice2 []int = arr[2:3]
	slice2[0] = 7
	println(arr[0], arr[1], arr[2], slice[1], slice2[0])
}
