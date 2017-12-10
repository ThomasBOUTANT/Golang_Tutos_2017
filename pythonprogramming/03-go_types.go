package main

import "fmt"


func add(x float64, y float64) float64 {
	return x+y
}

func multiple(a,b string) (string,string) {
	return a,b
}

func main() {
	var num1 float64 = 5.6
	var num2 float64 = 9.5

    fmt.Println("sum of", num1, "and", num2, "is", add(num1,num2))

	w1,w2 := multiple("Thomas","My name is")
	fmt.Println(w2,w1)
}