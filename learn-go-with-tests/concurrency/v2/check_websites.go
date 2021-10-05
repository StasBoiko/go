package main

import (
	"time"
	"fmt"
)

// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map.
// of urls to the result of checking each url with the WebsiteChecker function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			fmt.Println(u)
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}

	func main () {
		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}
	
		// want := map[string]bool{
		// 	"http://google.com":          true,
		// 	"http://blog.gypsydave5.com": true,
		// 	"waat://furhurterwe.geds":    false,
		// }
	
		CheckWebsites(mockWebsiteChecker, websites)
	
	}