package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	c := make(chan int)
	wg.Add(1)
	go producer(&wg, c)
	wg.Add(1)
	go reciver(&wg, c)
	wg.Wait()
}

func producer(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		c <- i
	}

	close(c)
}

func reciver(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for i := range c {
		fmt.Println("c:", i)
	}

}
