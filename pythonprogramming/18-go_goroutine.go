package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i:=0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}	
}


func main1() {
	fmt.Println("*** Main 1 ***")
	go say("1 - Hey")
	say("1 - there")
}


func main2() {
	fmt.Println("*** Main 2 ***")
	say("2.Hey")
	go say("2.there")
}

func main3() {
	fmt.Println("*** Main 3 ***")
	go say("#3 Hey")
	go say("#3 there")
}

func main4() {
	fmt.Println("*** Main 4 ***")
	go say("4 Hey")
	go say("4 there")
	time.Sleep(1000*time.Millisecond)
}

func main() {
	main1()
	main2()
	main3()
	main4()
}
