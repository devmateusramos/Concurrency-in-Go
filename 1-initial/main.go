package main

import (
	"fmt"
	"sync"
)

func printSomething(phrase string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(phrase)
}

func main() {
	var waitgroup sync.WaitGroup

	words := []string{
		"primeira palavra",
		"secunda palavra",
		"terceira palavra",
		"quarta palavra",
		"quinta palavra",
		"sexta palavra",
		"sétima palavra",
		"oitava palavra",
		"nona palavra",
	}

	waitgroup.Add(len(words))

	for _, word := range words {
		go printSomething(word, &waitgroup)
	}
	waitgroup.Wait()
	waitgroup.Add(1)
	printSomething("This is the second thing to be printed!", &waitgroup)
}

// First solution and the worst

//import (
//"fmt"
//"time"
//)
//
//func printSomething(phrase string) {
//	fmt.Println(phrase)
//}
//
//func main() {
//	go printSomething("This is the first thing to be printed!")
//	time.Sleep(time.Second)
//	printSomething("This is the second thing to be printed!")
//}

// Second solution wait groups

//import (
//"fmt"
//"sync"
//)
//
//func printSomething(phrase string, wg *sync.WaitGroup) {
//	defer wg.Done()
//	fmt.Println(phrase)
//}
//
//func main() {
//	var waitgroup sync.WaitGroup
//
//	waitgroup.Add(1)
//	go printSomething("This is the first thing to be printed!", &waitgroup)
//	waitgroup.Wait()
//
//	waitgroup.Add(1)
//	printSomething("This is the second thing to be printed!", &waitgroup)
//}
