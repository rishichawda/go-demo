package examples

import "fmt"

func panicDeferAndRecover() (i int) {
	i = 0
	defer func() {
		i += 1
	}()
	fmt.Println("Startin..")
	recoverFunction()
	// This gets executed because the flow was recovered in the defer call using recover.
	fmt.Println("panicDeferAndRecover func completed.")
	return
}

func recoverFunction() {
	fmt.Println("Recover function started executing..")
	defer func ()  {
		fmt.Printf("Recovered value: %v\n", recover())
	}()
	// First this defer & then the recover call is done as defer follows LIFO queue
	defer fmt.Println("Called method panicked. Starting recovery")
	fmt.Println("Calling the function that will panic.")
	panickingFunction()
	// The next line isn't executed because the method call above panics & the control moves to defer.
	fmt.Println("recoverFunction successfully completed.")
}

func panickingFunction() {
	fmt.Println("This function will panic.")
	// The next line will be called because defer is executed no matter what.
	defer fmt.Println("Okay now it panicked.")
	panic("Okay panicking.")
	// The next line is unreachable because the method will panic in the line above.
	fmt.Println("panickingFunction was done panicking")
}
