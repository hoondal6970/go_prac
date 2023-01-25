package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [6]string{"sol", "da", "ho", "chan", "jeon", "sik"}
	for _, person := range people {
		go issexy(person, c)
	}
	fmt.Println("waiting for messages")
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func issexy(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	channel <- person + " is sexy"
}
