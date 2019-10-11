package producer

import (
	"fmt"
)

const count = 10

var edmsg chan string

func consumer(name string, ch <-chan int) {
	for i := range ch {
		fmt.Println("consumer-", name, ":", i)
	}
	edmsg <- name
}

func producer(name string, ch chan<- int) {
	for index := 0; index < count; index++ {
		ch <- index
		fmt.Println("producer--", name, ":", index)
	}
	close(ch)
}

func RunChCache() {
	edmsg = make(chan string, 3)
	ch := make(chan int, 10)
	go producer("producer", ch)
	go consumer("user 1", ch)
	go consumer("user 2", ch)
	go consumer("user 3", ch)

	for index := 0; index < 3; index++ {
		fmt.Println("consumerEnd--", <-edmsg)
	}
	close(edmsg)
}
