package main

import (
	"fmt"
	"time"
)

// note that the time.Timer in golang standard lib is only meant to be used as a "one shot" timer.
// It sends a message through a channel after a certain duration, and thats it. It can be stopped ahead
// of that moment in time but can't be resumed. Not useful as an actual timer but more of an event
// timer. (execute "blah blah" code after this duration. retry api call after this duration)
func Timer() {
	seconds := 10
	timer := time.NewTimer(time.Duration(seconds) * time.Second)

	ticker := time.NewTicker(1 * time.Second)

	fmt.Printf("Countdown %d seconds \n", seconds)

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("\n Time's up!")
				return
			case <-ticker.C:
				seconds--
				fmt.Printf("\nRemaining: %d seconds", seconds)
			}
		}
	}()

	<-timer.C
	ticker.Stop()
}
