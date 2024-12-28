package main

import "fmt"

// null is nil

func iloveCPP() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct { // struct
	X int
	Y int
	// usage: fmt.Println(Vertex{1, 2}); v.X = 10;
	// v2 = Vertex{X: 1}  // Y:0 is implicit
	// var p *Vertex; a := v.X // same for plain and pointer to struct
	// vec = &Vertex{1, 2} // has type *Vertex
}

func arrays() {
	var arr [5]int // non-resizable, just like C
	arr = [5]int{1, 2, 3, 4, 5}
	arr[0] = 42
	// slices, just like Python
	slice := arr[1:3] // arr[:] equals whole array
	println(len(slice))
	fmt.Println(slice)
	s := []struct { // creates array and then a slice referencing it
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(s)
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s)
}

func dynamicArrays() {
	// The make function allocates a zeroed array and returns a slice that refers to that array
	// b := make([]type, size, capacity)
	matrix := [][]string{
		[]string{"_", "_"},
		[]string{"_", "_"},
	}
	fmt.Println("Total size", len(matrix)*2)
	// slices have append function that makes them Vectors from C++
	s := make([]int, 2)
	s = append(s, 2, 3, 4)
	fmt.Println(s)
}

func iterate() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow { // like enumerate in Python, also can use _, i syntax
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// for i := range pow {} would only use INDEX
}

func maps() {
	// var m map[string]float64 // nil, values can't be added
	m := make(map[string]float64)
	m["key"] = 1.1
	fmt.Println(m["key"])

	type Vertex struct {
		Latitude, Longitude float64
	}
	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	println(m2)
	// or just
	var m3 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	delete(m3, "Bell Labs")
	elem, ok := m3["Bell Labs"]
	if ok {
		fmt.Println(elem)
	}
}

// functions could als obe used as values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
	// fmt.Println(compute(math.Pow))
}

// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables;
// in this sense the function is "bound" to the variables.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func useClosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 3; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func main() {
	arrays()
	dynamicArrays()
	maps()
	WordCount("hello   moto moto you are awesome")
	useClosure()
}

// ============ assignment 1 (slices) ============

// package main
// import "golang.org/x/tour/pic"
func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		res[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			res[y][x] = uint8(x * y)
		}
	}
	return res
}

// func main() {
// 	pic.Show(Pic)
// }

// =========== assignment 2 (wordcount) ==========

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else {
			res[word] += 1
			word = ""
		}
	}
	res[word] += 1
	delete(res, "")
	fmt.Println("WordCount:", res)
	return res
}
