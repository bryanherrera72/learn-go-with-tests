package sync

import (
	"sync"
	"testing"
)

type Counter struct{
	mu sync.Mutex // better to define it like this in practice
	// this is because we prefer the mutex be scoped more privately
	// DONT use the alternate "sync.Mutex" without declaring it a private variable.
	value int
}

func (c *Counter) Inc(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int{
	return c.value
}

func TestCounter(t *testing.T){
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T){
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		//pass ref for our mutex, not a copy
		assertCounter(t, &counter, 3)
	})
	//NOTE: We're using waitgroup to hold until all goroutines wrap up.
	t.Run("it runs safe concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i:= 0; i < wantedCount; i++{
			go func(){
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		//pass ref for our mutex, not a copy
		assertCounter(t, &counter, wantedCount)
	})
}
//NOTE: a mutex shouldn't be copied after its first use. 
// we'll sustain the ref to it here. 
func assertCounter(t testing.TB, got *Counter, want int){
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}