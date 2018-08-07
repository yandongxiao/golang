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
	fmt.Println("raw path: ", uri.RawPath)
	fmt.Println("raw query: ", uri.RawQuery)
	fmt.Println("opaque: ", uri.Opaque)
}
