package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 500)
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}

}
