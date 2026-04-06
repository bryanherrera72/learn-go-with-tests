package helloworld

import "testing"

//This test now supports 'sub tests'
func TestHello(t *testing.T){
	// subtest 1. Say hello to someone
	t.Run("saying hello to people", func(t *testing.T) { 
		got := Hello("Bryan", "English")
		want := "Hello, Bryan"
		assertCorrectMessage(t, got, want)
	})
	//subtest 2. Default to hello world.
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got := Hello("Baguette", "French")
		want := "Bonjour, Baguette"
		assertCorrectMessage(t, got, want)
	})
}

//refactored the repeated block
// we used t testing.TB instead of a '*' to accept an interface. This method 
// accepts either a Test, or a Benchmark in this case. 
func assertCorrectMessage( t testing.TB, got, want string){// Shorthand: got and want are of type string.
	t.Helper() // Tells the test suite. "This is a Helper"
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}