package helloworld

 const (
	//prefixes
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	// language selectors
	spanish = "Spanish"
	french = "French"
)


func Hello(name, language string) string{
	if name == "" {
		name = "World"
	}
	
	// ***Non Switch ***
	// if language == spanish {
	// 	greeting = spanishHelloPrefix
	// } else if language == french{
	// 	greeting = frenchHelloPrefix
	// }
	return greetingPrefix(language) + name
}

//The return type creates a prefix var in the scope of our method
// the var starts with the zero value of its type (this case a ""). 
//Prefix will appear in the godoc for function greetingPrefix 
//This function is also private, since first letter of the name is lowercase.


// greetingPrefix takes a language as input and returns the prefix corresponding
// to that language.
func greetingPrefix(language string) (prefix string) { 
	switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case french: 
			prefix = frenchHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return
}

