package main

import "fmt"

func main() {
	x := 15
	a := &x         		// point to x  (memory address)

	fmt.Println("x :=", x, "and a := &x")

	fmt.Println("a =", a)   // prints out the mem addr.
	fmt.Println("*a =", *a) // read a through the pointer, so this will print out a value (15 in this case)
	
	*a = 5          		// sets the value pointed at to 5, which means x is modified (since x is stored at the mem addr)
	fmt.Println("*a :=", *a)
	fmt.Println("x =",x) 	// see the new value of x

	*a = *a**a      		// what is the value of x now? 
	fmt.Println("*a := *a * *a")
	fmt.Println("x =",x)	// prints a value
	fmt.Println("*a =",*a) 	// prints a value
	fmt.Println("&x =",&x) 	// prints a memory address
	fmt.Println("a =",a) 	// prints a memory address      
}