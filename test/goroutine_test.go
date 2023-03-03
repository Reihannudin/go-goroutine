package test

import (
	"fmt"
	"goroutine/index"
	"testing"
	"time"
)

func TestGoroutineHelloWorld(t *testing.T) {
	go index.RunHelloWorld()
	fmt.Println("Ups..")

	time.Sleep(10 * time.Second)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go index.DisplayManyGoroutine(i)
	}

	time.Sleep(10 * time.Second)
}
