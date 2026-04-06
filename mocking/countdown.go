package main

import (
	"fmt"
	"io"
	"time"
)

const finalWord =  "Go!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct{
	Calls int
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}


func(s *SpyCountdownOperations) Sleep(){
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error){
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct{
	duration time.Duration
	sleep func(time.Duration)
}

type SpyTime struct{
	durationSlept time.Duration
}
func (c *ConfigurableSleeper) Sleep(){
	c.sleep(c.duration)
}


func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper){
	for i := countdownStart; i > 0; i--{
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}