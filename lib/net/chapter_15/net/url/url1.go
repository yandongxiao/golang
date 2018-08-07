package main

import "net/url"
import "fmt"

func main() {
	uri, _ := url.Parse("https://www.google.com/search?q=https+http#d s")

	fmt.Println(uri)
	fmt.Println("scheme: ", uri.Scheme)

	if uri.User != nil {
		fmt.Println(uri.User)
	}

	fmt.Println("host: ", uri.Host)
	fmt.Println("path: ", uri.Path)
	fmt.Println("query: ", uri.Query())
	fmt.Println("Fragment: ", uri.Fragment)
}
