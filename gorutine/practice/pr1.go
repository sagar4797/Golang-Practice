package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go Generator(ch)
	time.Sleep(3)
	go Reciver(ch)
}

func Generator(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
	}

}
func Reciver(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
