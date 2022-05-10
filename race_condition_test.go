package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

//Bahaya di concurency (goroutine) bisa terjadi race condition
//Dimana beberapa goroutine yang berbeda mengakses variable saya yang di share
func TestReceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter ke ", x)

}
