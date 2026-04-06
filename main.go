package main

import (
	"os"

	"b.lang/learn_go_with_tests/mocking"
)


func main(){
	sleeper := &mocking.DefaultSleeper{}

	mocking.Countdown(os.Stdout, sleeper)
}

