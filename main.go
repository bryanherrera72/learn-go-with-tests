package main


func main(){
	//This throws an error since the struct literals are private 
	// and only assignable within the same package. Uncomment to see error
	// sleeper := &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}
	// mocking.Countdown(os.Stdout, sleeper)
}

