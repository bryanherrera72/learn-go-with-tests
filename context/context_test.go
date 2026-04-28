package context

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Store interface{
	Fetch(ctx context.Context) (string, error)
}
type SpyStore struct{
	response string
	t *testing.T
}
func (s *SpyStore) Fetch(ctx context.Context) (string, error){
	data := make(chan string, 1)

	go func(){
		var result string
		for _, c := range s.response{
			select{
			case <- ctx.Done():
				log.Println("spy store got cancelled")
				return
			default: 
				//simulate a slow process. building a string 
				//char by char
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result // write the completed string to the data channel
	}()

	select {
	case <- ctx.Done()://cancel was forwarded to this functions  main goroutine
		// propogate error upwards if the
		return "", ctx.Err()
	case res := <- data: //no errors, go routine finished work return channel 
		return res, nil
	}
}

type SpyResponseWriter struct{
	written bool
}

// These three methods implement the http.ResponseWriter interface
func (s *SpyResponseWriter) Header() http.Header{
	s.written = true
	return nil
}

func(s *SpyResponseWriter) Write([] byte) (int, error){
	s.written = true
	return 0, errors.New("not implemented")
}

func(s *SpyResponseWriter) WriteHeader(statusCode int){
	s.written = true
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		data, err := store.Fetch(r.Context())
		
		if err != nil{
			return // todo: log error
		}
		fmt.Fprint(w, data)
	}
}

func TestServer(t *testing.T){
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
	
		svr.ServeHTTP(response, request)
	
		if response.Body.String() != data{
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)// remember this runs fetch which will run for 100 
		//ms in our "SpyStore"

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		//cancel before 100 ms
		time.AfterFunc(5 * time.Millisecond, cancel) // here we're cancelling 
		request = request.WithContext(cancellingCtx) // we need to add the new cancelled context 

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written{
			t.Error("a response should not have been written")
		}
	})
}

