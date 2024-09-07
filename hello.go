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
	fmt.Println("Hi", myName)               // Joe Bloggs
	fmt.Print("Hi ", myName, "\n")          // Joe Bloggs
	fmt.Println("pi = ", PI)                // pi =  3.14159265358
	fmt.Println(a, b, c, d)                 // mars 42 true 21
	fmt.Println(nada1, nada2, nada3, nada4) // 0 0 false
	fmt.Println(e, f)                       // oops false
	// printf formatting verbs (type and value)
	fmt.Printf("b is of type %T and has value %v\n", b, b)
	fmt.Printf("%d\n", d)       // %d int
	fmt.Printf("%4dOops\n", d)  //    int right justified by 4
	fmt.Printf("%-4dOops\n", d) //    int left justified by 4
	fmt.Printf("%+d\n", d)      //    int always show sign
	fmt.Printf("%s\n", a)       // %s string
	fmt.Printf("%q\n", a)       // 	  string double quoted
	fmt.Printf("%8sOops\n", a)  // 	  string right justified
	fmt.Printf("%-8sOops\n", a) // 	  string left justified
	fmt.Printf("%f\n", PI)      // %f floats
	fmt.Printf("%.2f\n", PI)    //    float precision 2
	fmt.Printf("%4.2f\n", PI)   // 	  float precision 2 width 4

	// value (%v) formatting prints out anything, even arrays and slices
	var myArr = [5]int{1000, 1001, 1002, 1003, 1004}
	fmt.Printf("int is: %v, string is: %v, float is %v, array is %v\n", 100, "hello, world!", 3.14, myArr)

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
	// size of array will be auto determined based on number of elements provided during initialization
	var cars = [...]string{"bmw", "ford", "peugeot", "ferrari"}
	flts := [...]float32{1.1, 2.2, 3.3, 4}
	fmt.Println(nos)								// [101 201 301 401]	
	fmt.Println(drinks)							    // [coke pepsi sprite fanta]
	fmt.Println("cars[0] = ", cars[0])				// cars[0] =  bmw
	fmt.Println(flts)								// [1.1 2.2 3.3 4]
	fmt.Println("len of array cars is:", len(cars)) // len of array cars is: 4

	// empty slice (dynamic)
	sl1 := []int{}
	fmt.Println("slice len is", len(sl1))		// slice len is 0
	/* capacity ("cap") of underlying array in the slice
	   if more underlying array elements needed, Go will copy
	   to a new underlying array with (usually) twice the capacity */
	fmt.Println("slice cap is", cap(sl1))		// slice cap is 0
	fmt.Println(sl1)							// []	
	// non-empty slice (dynamic)
	sl2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println("slice len is", len(sl2))		// slice len is 4
	fmt.Println("slice cap is", cap(sl2))		// slice cap is 4
	fmt.Println(sl2)							// [Go Slices Are Powerful]	
	sl2 = append(sl2, "Growwwwing")
	fmt.Println("slice len is", len(sl2))		// slice len is 5
	fmt.Println("slice cap is", cap(sl2))		// slice cap is 8
	fmt.Println(sl2)							// [Go Slices Are Powerful Growwwwing]
	// create slice from array
	testArr := [3]int{1000, 1001, 1002}
	testSlice := testArr[0:2]
	fmt.Println("len testSlice", len(testSlice))	// len testSlice 2
	fmt.Println("cap testSlice", cap(testSlice))	// cap testSlice 3
	fmt.Println(testSlice)							// [1000 1001]
	fmt.Printf("testSlice: %v\n", testSlice)		// testSlice: [1000 1001]
}
