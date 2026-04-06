package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){
	got := Repeat("A", 6)
	want := "AAAAAA"
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
func ExampleRepeat(){
	str := "B"
	count := 5
	fmt.Println(Repeat(str,count))
	// Output: BBBBB
}

func BenchmarkRepeat(b *testing.B){
	//structured like: 
	//...setup...
	for b.Loop() {
		//...code to measure...
		Repeat("a", 6)
	}
	//...cleanup...
}

