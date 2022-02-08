package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quits chan struct{}) {
	//switch case
	//select for channel also async
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("recieved")
		case <-quits:
			fmt.Println("quit")
			os.Exit(0)

		}
	}
}
func main() {
	c := make(chan int) //unbuffered or buffered channel
	quits := make(chan struct{})

	go func() {
		c <- 1
		//quits <- struct{}{}
	}()
	go Select(c, quits)
	select {} //blocking code for show outputs
}
