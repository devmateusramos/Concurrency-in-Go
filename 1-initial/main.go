package main

import (
	"fmt"
	"time"
)

func printSomething(phrase string) {
	fmt.Println(phrase)
}

func main() {
	go printSomething("This is the first thing to be printed!")
	time.Sleep(time.Second)
	printSomething("This is the second thing to be printed!")
}
