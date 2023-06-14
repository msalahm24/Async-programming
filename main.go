package main

import (
	"fmt"
	"net/http"
	"sync"
	
)

func main(){
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	var wg sync.WaitGroup
	for _,url :=range urls{
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			checkUrl(url)
		}(url)
	}
	wg.Wait()
}

func checkUrl(url string){
	_,err := http.Get(url)

	if err != nil{
		fmt.Println("this site is down")
	}
	fmt.Println("this site is up")
}