package main

import "fmt"

func foo(c chan int, someValue int) {
    c <- someValue * 5
    fmt.Println("someValue = ", someValue)
}

func main() {
    fooVal := make(chan int)
    go foo(fooVal, 5)
    go foo(fooVal, 3)
    v1 := <-fooVal
    v2 := <-fooVal
    fmt.Println(v1, v2)


    go foo(fooVal, 2)
    go foo(fooVal, 6)
    v3 := <-fooVal
    v4 := <-fooVal
    fmt.Println(v3, v4)
}