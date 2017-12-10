package main

// https://golangbot.com/learn-golang-series/

import (  
    "fmt"
    "math"  // variables_operations() -> Min(), Max()
    "unsafe" // variables_types() -> Sizeof()
)

func main() {
	//variables_declaration()	
	//variables_operation()
	variables_types()
	variables_constants()
}

func variables_declaration() {
	fmt.Println("\n*** Variables Declaration ***\n")

	// Declaring a single variable
	var name string
	name = "Thomas"
    fmt.Println("   my name is", name)

    // Declaring a variable with initial value
    var age int = 23
    fmt.Println("   my age is", age)

    // Multiple variable declaration
    var width, height int = 100, 50
    fmt.Println("   width is", width, "height is", height)

    // Declare variables belonging to different types in a single statement
    var (
        languages = "golang"
        date int = 2017
    )
    fmt.Println("   I've begun to learn", languages, "in", date)

    // Short hand declaration
    texte, entier := "string", 4 
    fmt.Println("   we have a", texte, "and the number", entier)

    a, b := 20, 30 // declare variables a and b
    fmt.Println("   a is", a, "b is", b)
    b, c := 40, 50 // b is already declared but c is new
    fmt.Println("   b is", b, "c is", c)
    b, c = 80, 90 // assign new values to already declared variables b and c
    fmt.Println("   changed b is", b, "c is", c)
}

func variables_operation() {
	fmt.Println("\n*** Variables Operation ***\n")

	// Min(a,b)
	a, b := 145.8, 543.8
    c := math.Min(a, b)
    fmt.Println("   minimum value is ", c)

    // Max(a,b)
    d := math.Max(a, b)
    fmt.Println("   maximum value is ", d)
}

func variables_types(){
	fmt.Println("\n*** Variables Types ***\n")

	// Bool
	a0 := true
    b0 := false
    fmt.Println("\n   a0:", a0, "b0:", b0)
    c := a0 && b0
    fmt.Println("   c:", c)
    d := a0 || b0
    fmt.Println("   d:", d)

    // Signed integers
    var a1 int = 89
    b1 := 95
    fmt.Println("\n   value of a1 is", a1, "and b1 is", b1)
    fmt.Printf("\n   type of a1 is %T, size of a1 is %d bytes", a1, unsafe.Sizeof(a1)) //type and size of a1
    fmt.Printf("\n   type of b1 is %T, size of b1 is %d bytes", b1, unsafe.Sizeof(b1)) //type and size of b2

    // Unsigned integers
    var a2 uint = 89
    b2 := 95
    fmt.Println("\n   value of a2 is", a2, "and b2 is", b2)
    fmt.Printf("\n   type of a2 is %T, size of a2 is %d bytes", a2, unsafe.Sizeof(a2)) //type and size of a2
    fmt.Printf("\n   type of b2 is %T, size of b2 is %d bytes", b2, unsafe.Sizeof(b2)) //type and size of b2


    // Floating point types
    a3, b3 := 5.67, 8.97
    fmt.Printf("\n\n   type of a3 %T b3 %T\n", a3, b3)
    sum := a3 + b3
    diff := a3 - b3
    fmt.Println("   sum", sum, "diff", diff)

    no1, no2 := 56, 89
    fmt.Println("   sum", no1+no2, "diff", no1-no2)


    // Complex type
    c1 := complex(5, 7)
    c2 := 8 + 27i
    cadd := c1 + c2
    fmt.Println("\n   sum:", cadd)
    cmul := c1 * c2
    fmt.Println("   product:", cmul)

    // String type
    first := "Thomas"
    last := "BOUTANT"
    name := first +" "+ last
    fmt.Println("   My name is", name)

    // Type Conversion
    i := 55      //int
    j := 67.8    //float64
    sum2 := i + int(j) //j is converted to int
    fmt.Println("\n   sum =", sum2)

    i2 := 10
    var j2 float64 = float64(i2) //this statement will not work without explicit conversion
    fmt.Println("   j2", j2)
}

func variables_constants(){
	fmt.Println("\n*** Variables Constants ***\n")

}