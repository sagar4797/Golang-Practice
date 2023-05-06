package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	wg := sync.WaitGroup{}
// 	c := make(chan int)
// 	wg.Add(1)
// 	go producer(c, &wg)
// 	wg.Add(1)
// 	go reciver(c, &wg)
// 	wg.Wait()
// }

// func producer(c chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 	}
// 	close(c) // Close the channel after sending all values
// }

// func reciver(c chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

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
