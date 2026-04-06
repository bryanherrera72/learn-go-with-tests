package iteration

import "strings"

func main(){
	
}

//Repeat accepts a string 'char' and int 'count'
// The function repeats 'char', x amount of times where x = 'count'
func Repeat(char string, count int) string{
	var repeated strings.Builder
	for i:= 0; i< count; i++{
		repeated.WriteString(char)
	}

	return repeated.String()
}