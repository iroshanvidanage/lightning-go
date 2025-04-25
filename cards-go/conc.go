package main

import (
	"fmt"
	"net/http"
	"time"
)

func statusChecker() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://go.dev",
	}

	c := make(chan string)

	for _, link := range links {
		// placing go before will execute a new routine
		go checkLink(link, c)
		// keeping the print statement will only executes each routine
		// sequencially cause the main routine will waits for the msg to come.
		// fmt.Println(<-c)
	}

	// This is a blocking signal
	// fmt.Println(<-c)
	// At this line it will only outputs the first child routines data and exits the call.
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// adding multiple lines will give all msgs but it's not efficient.

	for l := range c {
		// fmt.Println(<-c)
		// instead of a function we can set a function literal (lambda)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
		// time.Sleep(5 * time.Second)
	}

}

func checkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- link + " might be down!"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- link + " is up!"
	c <- link
}
