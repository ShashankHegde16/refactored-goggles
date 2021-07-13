package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com/",
		"https://udemy.com/",
		"https://www.linkedin.com/",
		"https://stackoverflow.com/",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLinks(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinks(link, c)
		}(l)

	}
}

func checkLinks(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link

}
