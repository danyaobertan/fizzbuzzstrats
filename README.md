# FizzBuzz Application
This is a simple FizzBuzz application written in Go that demonstrates the use of the Strategy design pattern. The application generates FizzBuzz answers for numbers from 1 to N and provides a flexible way to output the results either to the console or to any other destination using the Strategy pattern.

## Project Structure
The project is organized as follows:

- `fizzbuzz` : The fizzbuzz package contains the main logic and interfaces for generating and writing FizzBuzz answers.
  - `fizzbuzz.go` : Contains the FizzBuzzWriter interface that defines the behavior for generating and writing FizzBuzz answers.
  - `fizzbuzz_basic.go` : Implements a basic FizzBuzz generator, `BasicFizzBuzzer`, that adheres to the `FizzBuzzWriter` interface. It generates FizzBuzz answers for a given integer and writes the results to the console.
- `main.go` : The `main.go` file contains the entry point of the application. It creates an instance of `BasicFizzBuzzer`, uses the `FizzBuzzContext` to generate and print FizzBuzz answers.
- `go.mod` : The `go.mod` file contains the module definition for the project.

## Strategy Pattern Usage
The Strategy pattern is utilized in this application to enable flexible generation and output of FizzBuzz answers. The `FizzBuzzWriter` interface defines two methods:

1. `FizzBuzz(i int) []byte` : This method is responsible for generating the FizzBuzz answer for a given integer i. Different implementations of `FizzBuzzWriter` can provide different FizzBuzz generation rules.
2. `Write(result [][]byte) error` : This method is used to write the FizzBuzz results. By implementing this method differently, it is possible to output the results to various destinations, such as the console, a file, a network stream, etc.

The `BasicFizzBuzzer` implements the FizzBuzzWriter interface, providing a simple FizzBuzz generation rule and writing the results to the console.

## Usage
To run the application and see the FizzBuzz answers printed to the console, execute the following command:

```bash
go run main.go
```

The application will generate FizzBuzz answers for the numbers from `1` to `n` (where `n` is set to `30`) using the BasicFizzBuzzer. The results will be printed to the console in the following format:

```
-  answer[1] == 1    
-  answer[2] == 2    
-  answer[3] == Fizz 
-  answer[4] == 4    
-  answer[5] == Buzz 
-  answer[6] == Fizz 
-  answer[7] == 7    
-  answer[8] == 8    
-  answer[9] == Fizz 
-  answer[10] == Buzz
-  answer[11] == 11
-  answer[12] == Fizz
-  answer[13] == 13
-  answer[14] == 14
-  answer[15] == FizzBuzz
-  answer[16] == 16
-  answer[17] == 17
-  answer[18] == Fizz
-  answer[19] == 19
-  answer[20] == Buzz
-  answer[21] == Fizz
-  answer[22] == 22
-  answer[23] == 23
-  answer[24] == Fizz
-  answer[25] == Buzz
-  answer[26] == 26
-  answer[27] == Fizz
-  answer[28] == 28
-  answer[29] == 29
-  answer[30] == FizzBuzz
```