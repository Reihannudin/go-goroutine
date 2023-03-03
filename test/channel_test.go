package test

import (
	"fmt"
	"goroutine/index"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Jeon Jungkook"
		fmt.Println("Finished sending data")
	}()

	data := <-channel
	fmt.Println(data)
	fmt.Println("Successfully receive data")

	time.Sleep(5 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go index.GiveMeResponseChannel(channel)

	data := <-channel
	fmt.Println(data)
	fmt.Println("Successfully receive data")

	time.Sleep(5 * time.Second)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go index.OnlyIn(channel)
	go index.OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 5)
	defer close(channel)

	go func() {
		channel <- "RM"
		channel <- "Suga"
		channel <- "J-Hope"
		channel <- "Jimin"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Finished")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go index.GiveMeResponseChannel(channel1)
	go index.GiveMeResponseChannel(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go index.GiveMeResponseChannel(channel1)
	go index.GiveMeResponseChannel(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Waiting Data")
		}
		if counter == 2 {
			break
		}
	}

}
