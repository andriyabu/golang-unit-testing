package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Andri")

	if result != "Hello Andri" {
		// panic("Test result Fail")
		t.Fail()
	}
	// still execute
	fmt.Println("Test Helloworld is done")
}

func TestHelloWordAbel(t *testing.T) {
	result := HelloWorld("Abe")

	if result != "Hello Abel" {
		t.FailNow()
	}
	// not execute
	fmt.Println("Test Hello World Abel is done")
}
