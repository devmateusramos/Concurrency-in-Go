package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestUpdateMessage(t *testing.T) {
	var wgTest sync.WaitGroup

	wgTest.Add(1)
	go updateMessage("epsilon", &wgTest)
	wgTest.Wait()

	if msg != "epsilon" {
		t.Error("Expected to find epsilon, but it is not there")
	}
}

func TestPrintMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	stdOut = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Error("Expected to find epsilon, but it is not there")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, first go!") {
		t.Errorf("Expected to find 'Hello, first go!', but it is not there : %s", output)
	}
	if !strings.Contains(output, "Hello, second go!") {
		t.Error("Expected to find 'Hello, second go!', but it is not there")
	}
	if !strings.Contains(output, "Hello, third go!") {
		t.Error("Expected to find 'Hello, third go!', but it is not there")
	}
}
