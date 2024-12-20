package main

// programm starts from package named main

import (
	"fmt"
	"math/rand"
	// some other import
)

// multiple imports possible but not recoimmended

// In Go, a name is exported if it begins with a capital letter
// When importing a package, you can refer only to its exported names

const Pi float64 = 10

func getRandomNumber(rangeStart, rangeEnd int) (value int) { // Notice that the type comes after the variable name (both int here)
	value = rangeStart + (rand.Int() % rangeEnd)
	// have int and uint for C lovers
	return
	// returns value. "Naked" return
}

func swap(x, y string) (string, string) {
	return y, x
}

func declareWithVar() int {
	var i, j, k int = 1, 2, 3 // huh?
	var string, or_int = "wtf", 42
	fmt.Println(string, or_int)
	res := i + j + k
	return res
} // := assignment not available outside of functions

func typeConversion() float64 {
	i := 42 // types are inferred
	f := float64(i)
	g := 0.867 + 0.5i
	fmt.Println("Meshuggah!! Go has complex vars!", g)
	return f
}

func main() {
	str := "Hello, world"
	fmt.Println(str)
	random := getRandomNumber(1, 1000)
	fmt.Println("Random of the day:", random)
	a, b := swap("b", "a") // var can be ommited when ":=" used
	fmt.Println("Swapped values:", a, b)
	fmt.Printf("Type: '%T' Value: '%v'\n", a, a) // get type example
}

// To execute:
// go run 1_app/file.go
// cd 1_app; go build file.go # -> executable
