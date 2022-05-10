package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Indra")
	pool.Put("Bayu")
	pool.Put("Sudirman")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

}
