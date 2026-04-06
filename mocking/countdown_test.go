package mocking

import (
	"bytes"
	"testing"
)

//works well however, this takes 3 seconds to complete. Not ideal if it begins
// to scale
// "slow tests ruin developer productivity". Investcloud loves slow tests.
func TestCountdown(t *testing.T){
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3{
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}