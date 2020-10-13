package examples

import "fmt"

func GoToExample()  {
	i := 0
	Start:
	if i >= 5 {
		fmt.Print(i)
		fmt.Print("\n")
		goto End
		} else {
			fmt.Print(i, ",")
			i += 1
			goto Start
		}
	End:
}

func BreakExample() {
	a := 0
	outerLoop:
	for {
		for i:=0;i<5;i++ {
			a = (a * i) + i
			fmt.Println(a)
			if a > 10 {
				break outerLoop
			}
		}
	}
}

func ContinueExample() {
	a := 0
	outerLoop:
	for j:=0;j<=2;j++ {
		for i:=0;i<5;i++ {
			a = (a * i) + i
			fmt.Println(a)
			if a > 10 {
				continue outerLoop
			}
		}
	}
}