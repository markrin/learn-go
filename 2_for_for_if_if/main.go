package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 3; i++ { // unfortunately { can't be on other line :(((((
		fmt.Println("FOR IS THE ONLY LOOP IN GO")
	}
	square := 1
	for square < 20 { // u can use full C style for(; a<3;) also
		square += square
	}
	if a := 10; square > 40 { // { } are required
		fmt.Println("How?")
	} else {
		println("We have a at home '%s'", a)
	}
	deferIsStrange() // can be declared anywhere
}

func whichOS() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux. The best.")
		// No need for break
	case "darwin":
		fmt.Println("OS X")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func ifEllseChain(input int) {
	switch {
	case input > 0:
		fmt.Println("Good")
	case input < 0:
		fmt.Println("Bad")
	default:
		fmt.Println("Ugly")
	}
}

// https://go.dev/blog/defer-panic-and-recover
func deferIsStrange() {
	for i := 1; i < 11; i++ {
		defer println(i)
	}
	// The list of saved calls is executed after the surrounding function returns
	println("Defer puts functions in a stack")
}
