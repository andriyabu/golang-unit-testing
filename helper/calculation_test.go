package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
