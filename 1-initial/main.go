package main

import "fmt"

func printSomething(phrase string) {
	fmt.Println(phrase)
}

func main() {
	go printSomething("This is the first thing to be printed!")
	printSomething("This is the second thing to be printed!")
}
