package dependency

import (
	"fmt"
	"io"
	"net/http"
)

// func Greet(writer *bytes.Buffer, name string){
func Greet(writer io.Writer, name string){	
	fmt.Fprintf(writer, "Hello, %s\n", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request){
	Greet(w, "world")
}

