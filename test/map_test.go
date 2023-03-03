package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	var data sync.Map
	var addToMap = func(value int) {
		data.Store(value, value)
	}

	for i := 0; i < 100; i++ {
		go addToMap(i)
	}

	time.Sleep(3 * time.Second)
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
