package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(100 * time.Millisecond)
	}

	channel <- "Done"
}

func main() {
	channel := make(chan string)

	// go printMessage("Go is great!")
	go printMessage("Go is awesome!", channel)
	// printMessage("To be or not to be")

	response := <-channel

	println(response)

}
