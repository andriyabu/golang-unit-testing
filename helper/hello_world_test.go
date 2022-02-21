package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Andri")

	if result != "Hello Andri" {
		// panic("Test result Fail")
		// t.Fail()

		t.Error("Result must be 'Hello Andri'")
	}
	// still execute
	fmt.Println("Test Helloworld is done")
}

func TestHelloWorldAbel(t *testing.T) {
	result := HelloWorld("Abel")

	if result != "Hello Abel" {
		// t.FailNow()
		t.Fatal("Result must be 'Hello Abel'")
	}
	// not execute
	fmt.Println("Test Hello World Abel is done")
}
