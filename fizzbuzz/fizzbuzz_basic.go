package fizzbuzz

import (
	"fmt"
	"strconv"
)

// BasicFizzBuzzer is a basic FizzBuzz generator that adheres to the FizzBuzzer interface
type BasicFizzBuzzer struct{}

// FizzBuzz generates the FizzBuzz answer for a given integer i
func (b BasicFizzBuzzer) FizzBuzz(i int) []byte {
	switch {
	case i%15 == 0:
		return []byte("FizzBuzz")
	case i%5 == 0:
		return []byte("Buzz")
	case i%3 == 0:
		return []byte("Fizz")
	default:
		return []byte(strconv.Itoa(i))
	}
}

// WriteFizzBuzz writes the FizzBuzz answer to an io.Writer (console)
func (b BasicFizzBuzzer) Write(result [][]byte) error {
	fmt.Printf("\nBasicFizzBuzzer:\n\n")
	for k, v := range result {
		fmt.Printf("-  answer[%d] == %s\n", k+1, string(v))
	}
	return nil
}
