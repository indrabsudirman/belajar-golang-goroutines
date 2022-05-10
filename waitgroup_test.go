package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second) //2.12.20

}
