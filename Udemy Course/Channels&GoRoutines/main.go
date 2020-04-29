package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// in the current state of the program it does not constantly check for the staus of the links

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	// this is a linear operation so it takes time , a parrallel operational approach will be better
	// main routines have higher priority than child routines therefore the program ends before any child routines finish their http requests
	// in order to avoid this we need to use channels
	c := make(chan string)
	for _, link := range links {
		go httpRequest(link, c)

	}

	// if only 1 print statement was here the program would stop after recieving the first channel
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	//this for loop iterates throught the loop of links for a number of channels equal to n elements
	//instead we create an infinite loop but will only be called as we ping the websites
	for l := range c {

		// //to prevent a blockage of flow we wrap the sleep code in a function literal(equivalent to an anonamous function in java)
		// go func() {
		// 	time.Sleep(3 * time.Second)
		// 	//with just these 2 lines of code we run into an error where after the first set of requests the for loop only recieves the facebook channel this is because go is a pass by value language and comes from l
		// 	httpRequest(l, c)

		// }()

		//instead we do this
		go func(link string) {
			time.Sleep(3 * time.Second)
			httpRequest(link, c)

		}(l)
	}
}

func httpRequest(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
