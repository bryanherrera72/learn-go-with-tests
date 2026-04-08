package concurrency

type WebsiteChecker func(string) bool

type result struct{
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string] bool{
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(){ // this anon function can use variables from the surrounding scope. 
			resultChannel <- result{url, wc(url)} // 'send' statement from the go routine. (note the channel is on left)
		}() //these parens are here because we're immediately calling the anon function.
	}

	// we're expecting a result for each url here (3 in the test case). so we iterate through them and receive 
	// each result from the channel.
	for i:= 0; i<len(urls); i++ {
		r := <- resultChannel // 'receive' expression
		results[r.string] = r.bool
	}

	return results
}