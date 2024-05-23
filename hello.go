package main

import (
	"fmt"
)

func main() {
	// variables
	var myName string = "Joe Bloggs"
	const PI float32 = 3.14159265358
	var a, b, c = "mars", 42, true
	d := 21
	// truthy falsey
	var (
		nada1 int
		nada2 string
		nada3 float32
		nada4 bool
	)
	e, f := "oops", false

	// snippet "fmp" "fmpl"
	fmt.Println("Hi", myName)
	fmt.Print("Hi ", myName, "\n")
	fmt.Println("pi = ", PI)
	fmt.Println(a, b, c, d)
	fmt.Println(nada1, nada2, nada3, nada4)
	fmt.Println(e, f)
	// printf formatting verbs (type and value)
	fmt.Printf("b is of type %T and has value %v\n", b, b)
	fmt.Printf("%d\n", d)       // %d ints
	fmt.Printf("%+d\n", d)      // always show sign
	fmt.Printf("%4dOops\n", d)  // right justified by 4
	fmt.Printf("%-4dOops\n", d) // left justified by 4
	fmt.Printf("%s\n", a)       // %s string
	fmt.Printf("%q\n", a)       // double Quoted string
	fmt.Printf("%8sOops\n", a)  // right justified
	fmt.Printf("%-8sOops\n", a) // left justified
	fmt.Printf("%f\n", PI)      // %f floats
	fmt.Printf("%.2f\n", PI)    // precision 2
	fmt.Printf("%4.2f\n", PI)   // precision 2 width 4

	// if statement
	if b > d {
		fmt.Println("b is greater than d")
	} else if b == d {
		fmt.Println("b == d")
	} else {
		fmt.Println("d greater than b")
	}

	// array
	var nos = [4]uint{101, 201, 301, 401}
	var drinks = [4]string{"coke", "pepsi", "sprite", "fanta"}
	var cars = [...]string{"bmw", "ford", "peugeot", "ferrari"}
	flts := [...]float32{1.1, 2.2, 3.3, 4}
	fmt.Println(nos)
	fmt.Println(drinks)
	fmt.Println("cars[0] = ", cars[0])
	fmt.Println(flts)
	fmt.Println("len of array cars is:", len(cars))

	// empty slice (dynamic)
	sl1 := []int{}
	fmt.Println("slice len is", len(sl1))
	/* capacity of underlying array in the slice
	   if more underlying array elements needed, Go will copy
	   to a new underlying array with (usually) twice the capacity */
	fmt.Println("slice cap is", cap(sl1))
	fmt.Println(sl1)
	// non-empty slice (dynamic)
	sl2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println("slice len is", len(sl2))
	fmt.Println("slice cap is", cap(sl2))
	fmt.Println(sl2)
	sl2 = append(sl2, "Growwwwing")
	fmt.Println("slice len is", len(sl2))
	fmt.Println("slice cap is", cap(sl2))
	fmt.Println(sl2)
	// create slice from array
	testArr := [3]int{1000, 1001, 1002}
	testSlice := testArr[0:2]
	fmt.Println("len testSlice", len(testSlice))
	fmt.Println("cap testSlice", cap(testSlice))
	fmt.Println(testSlice)
	fmt.Printf("testSlice: %v\n", testSlice)
}
