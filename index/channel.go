package index

import (
	"fmt"
	"time"
)

func GiveMeResponseChannel(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Kim Seokjin"
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Kim Taehyung"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}
