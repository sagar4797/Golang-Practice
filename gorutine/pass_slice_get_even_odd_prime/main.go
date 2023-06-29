package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	oddCh := make(chan int)
	evenCh := make(chan int)
	primeCh := make(chan int)
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(1)
	go ReturnEven(wg, evenCh, s)
	for evenVal := range evenCh {
		fmt.Println(evenVal)
	}
	wg.Add(1)
	go RetunOdd(wg, oddCh, s)
	for oddVal := range oddCh {
		fmt.Println(oddVal)
	}
	wg.Add(1)
	go RetunPrime(wg, primeCh, s)
	for primeVal := range primeCh {
		fmt.Println(primeVal)
	}
	wg.Wait()

}

func ReturnEven(wg *sync.WaitGroup, ch chan int, s []int) {
	defer wg.Done()
	for _, val := range s {
		if val%2 == 0 {
			ch <- val
		}
	}
	close(ch)
}

func RetunOdd(wg *sync.WaitGroup, ch chan int, s []int) {
	defer wg.Done()
	for _, val := range s {
		if val%2 != 0 {
			ch <- val
		}
	}
	close(ch)
}

func RetunPrime(wg *sync.WaitGroup, ch chan int, s []int) {
	defer wg.Done()
	flag := true
	for _, val := range s {
		for i := 2; i < val/2; i++ {
			if val%i == 0 {
				flag = false
			}
		}

		if flag {
			ch <- val
		}
		flag = true
	}
	close(ch)
}
