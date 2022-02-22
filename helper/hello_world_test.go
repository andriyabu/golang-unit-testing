package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		HelloWorld("Andri")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(andri)",
			request:  "andri",
			expected: "Hello andri",
		},
		{
			name:     "HelloWorld(yabu)",
			request:  "yabu",
			expected: "Hello yabu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTestHello(t *testing.T) {
	t.Run("assertion", func(t *testing.T) {
		result := HelloWorld("sayang")

		assert.Equal(t, "Hello sayang", result, "Result must be 'Hello sayang'")
		fmt.Println("Sub test with assertion done")
	})

	t.Run("require", func(t *testing.T) {
		result := HelloWorld("Andri")

		require.Equal(t, "Hello Andri", result, "Result must be 'Hello Andri'")
		fmt.Println("Sub test with require done")
	})
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("sayang")

	assert.Equal(t, "Hello sayang", result, "Result must be Hello sayang")
	fmt.Println("Test Hello world with Assertion done")
}

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
