package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	c := make(chan int)
	go producer(c, &wg)
	wg.Add(1)
	go retrive(c, &wg)
	wg.Add(1)

	wg.Wait()
}

func producer(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		c <- i
	}

	close(c)
}

func retrive(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range c {
		fmt.Println(i)
	}
}
