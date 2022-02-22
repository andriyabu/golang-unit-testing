package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		result := Add(3, 2)
		require.Equal(t, 5, result, "Resul must be 5")
		fmt.Println("Sub test Add is done")
	})

	t.Run("multiply", func(t *testing.T) {
		result := Multiply(3, 3)
		require.Equal(t, 9, result, "Result must be 9")
		fmt.Println("Sub test Multiply is done")
	})

	t.Run("devided", func(t *testing.T) {
		result := Devided(50, 5)
		require.Equal(t, 10, result, "Result must be 10")
		fmt.Println("Sub test Devided is done")
	})
}

func TestMain(m *testing.M) {
	//before testing
	fmt.Println("Show before execute testing")
	m.Run()

	//after testing
	fmt.Println("Show after execute testing")
}

func TestCalculationSkippAdd(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("can not run on Linux")
	}

	result := Add(2, 2)
	require.Equal(t, 4, result, "Result must be 4")
	fmt.Println("Test Skipp Add Calculation done")
}

func TestCalculationAdd(t *testing.T) {
	result := Add(5, 5)
	// assert.Equal() will execute testing.Fail()
	assert.Equal(t, 10, result, "Result must be 10")
	fmt.Println("Test Add Calculations Done")
}

func TestCalculationMultiply(t *testing.T) {
	result := Multiply(5, 5)

	// require.Equal() will execute testing.FailNow()
	require.Equal(t, 25, result, "Result must be 25")
	// this message willnot be execute
	fmt.Println("Test Multiply calculation Done")
}

func TestCalculationDivide(t *testing.T) {
	result := Devided(25, 5)

	require.Equal(t, 5, result, "Result must be 5")
	fmt.Println("Test Devided calculation Done")
}
