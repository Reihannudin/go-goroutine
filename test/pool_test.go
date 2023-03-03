package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	//var pool sync.Pool

	//automatic pool
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Andrian")
	pool.Put("Reihan")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	time.Sleep(3 * time.Second)
}
