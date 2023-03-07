package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 45, 32, 65, 23, 73, 76, 232}

	fmt.Println("s1", s1)
	s2 := s1[1:5]
	fmt.Println("s2", s2)

}

// func foo(c chan int, val int) {
// 	c <- val
// }
// func main() {
// 	mycha := make(chan int)
// 	go foo(mycha, 5)
// 	v1 := <-mycha
// 	fmt.Println(v1)
// }

// func test1() {
// 	fmt.Println("Inside the test1")
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("In test")
// 	}

// }

// func main() {
// 	mych := make(chan int)
// 	wg := &sync.WaitGroup{}
// 	mych <- 5
// 	wg.Add(1)
// 	go test(mych, wg)
// 	wg.Wait()
// }

// func test(mych chan int, wg *sync.WaitGroup) <-chan int {
// 	data := <-mych
// 	fmt.Println(data)
// 	return mych
// }

//1. -------
//printOutput
// defer fmt.Println("1")
// defer fmt.Println("2")
// defer fmt.Println("3")
// fmt.Println("4")
