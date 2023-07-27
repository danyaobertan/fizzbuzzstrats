package main

import (
	"github.com/danyaobertan/fizzbuzzstrats/fizzbuzz"
)

// n is the number of FizzBuzz answers to generate
const n = 30

func main() {
	// Create an instance of the BasicFizzBuzzer
	basicFizzBuzzer := fizzbuzz.BasicFizzBuzzer{}
	// Use the FizzBuzzContext to generate and print FizzBuzz answers using BasicFizzBuzzer
	fizzBuzzContext := fizzbuzz.NewContext(basicFizzBuzzer)
	fizzBuzzContext.FizzBuzz(n)
}
