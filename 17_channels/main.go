package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// ch := make(chan int)
	// for j := 0; j < 5; j++ {
	// 	wg.Add(2)
	// 	go func() {
	// 		i := <-ch
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()
	// 	go func() {
	// 		ch <- 42
	// 		wg.Done()
	// 	}()
	// }
	// wg.Add(2)
	// go func() {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	ch <- 27
	// 	wg.Done()
	// }()
	// go func() {
	// 	ch <- 42
	// 	fmt.Println(<-ch)
	// 	wg.Done()
	// }()
	// // receive only channel
	// go func(ch <- chan int) {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }(ch)
	// // send only channel
	// go func(ch chan <- int) {
	// 	ch <- 42
	// 	wg.Done()
	// }(ch)

	// buffered channel(not an ideal solution for this problem)
	// ch := make(chan int,50)
	// wg.Add(2)

	// go func(ch <- chan int) {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	i = <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }(ch)
	// go func(ch chan <- int) {
	// 	ch <- 42
	// 	ch <- 27
	// 	wg.Done()
	// }(ch)

	ch := make(chan int,50)
	wg.Add(2)

	go func(ch <- chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		ch <- 42
		ch <- 27
		// close the channel after sending all the data to avoid receiver deadlock
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

}
