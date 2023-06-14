package main

import (
	"fmt"
	"net/http"
)

type urlStatus struct {
	url    string
	status bool
}

func main() {
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	ch := make(chan urlStatus, len(urls))
	for _, url := range urls {
		checkUrl(url, ch)
	}

	results := make([]urlStatus, len(urls))

	for i,_ := range results {
		results[i] = <- ch
		if results[i].status{
			fmt.Println(results[i].url + " " + "is up")
		}
		if ! results[i].status{
			fmt.Println(results[i].url + " " + "is down")
		}
	}

}

func checkUrl(url string, ch chan urlStatus) {
	_, err := http.Get(url)

	if err != nil {
		ch <- urlStatus{url: url, status: false}
		return
	}
	ch <- urlStatus{url: url, status: true}
}
