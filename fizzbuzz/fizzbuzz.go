package fizzbuzz

// FizzBuzzWriter defines the behavior for generating FizzBuzz answers
type FizzBuzzWriter interface {
	FizzBuzz(i int) []byte       // Method to generate the FizzBuzz answer for a given integer i
	Write(result [][]byte) error // Method to write the FizzBuzz answer
}

// Context is the client structure that uses the FizzBuzzer interface
type Context struct {
	fizzBuzzWriter FizzBuzzWriter
}

// NewContext creates a new Context with a specific FizzBuzzer implementation
func NewContext(fbw FizzBuzzWriter) *Context {
	return &Context{
		fizzBuzzWriter: fbw,
	}
}

// FizzBuzz returns the FizzBuzz answers for numbers from 1 to n
func (c *Context) FizzBuzz(n int) [][]byte {
	// Initialize the result slice with length and capacity of n
	result := make([][]byte, n, n)
	for i := 1; i <= n; i++ {
		result[i-1] = c.fizzBuzzWriter.FizzBuzz(i)
	}
	err := c.fizzBuzzWriter.Write(result)
	if err != nil {
		return nil
	}
	return result
}
