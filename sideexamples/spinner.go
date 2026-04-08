package main

//This program shows a goroutine animation running concurrently with a fibonacci calculation
//
import (
	"fmt"
	"time"
)

func SpinnerEx(){
	//In this example, the spinner is running concurrently alongside the calculation.
	go spinner(100 * time.Millisecond)

	const n = 45
	//The "main" goroutine is running the fibonacci calculation. When it completes,
	// the program exits and the goroutine ''spinner'' gets resolved.
	fibN := fib(n) // slow version
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration){
	for{
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int{
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib( x - 2)
}
