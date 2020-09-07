package main

import (
	"fmt"
	// "runtime"
	"sync"
);

// to run a func using go routine add "go" infront of func calling...added a delay so that go routine can execute func.
// func main() {
// 	go sayHello();
// 	time.Sleep(100 * time.Millisecond)
// }

// func sayHello(){
// 	fmt.Println("Hello")
// }

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

// passing argument to a go routine is recommnended instead of using closures...as it may lead to race condition.
// func main(){
// 	var msg = "Hello"
// 	wg.Add(1)
// 	go func(msg string) {
// 		fmt.Println(msg)
// 		wg.Done()
// 	}(msg)
// 	msg = "Goodbye"
// 	wg.Wait()
// }


// // using only waitgroup
// func main()  {
// 	for i:=0 ; i <10 ; i++ {
// 		wg.Add(2)
// 		go sayHello()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello()  {
// 	fmt.Printf("Hello #%v\n",counter)
// 	wg.Done()
// }

// func increment()  {
// 	counter++
// 	wg.Done()
// }

// In this case it would execute in line but sometimes what would happen is read would execute twice before writing.

// func main(){
// 	// runtime.GOMAXPROCS(100) 
// 	for i := 0 ; i<10;i++{
// 		wg.Add(2)
// 		go sayHello()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello()  {
// 	m.RLock()
// 	fmt.Printf("Hello #%v\n",counter)
// 	m.RUnlock()
// 	wg.Done()
// }

// func increment()  {
// 	m.Lock()
// 	counter++
// 	m.Unlock()	
// 	wg.Done()
// }

func main(){
	for i := 0 ; i<10;i++{
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello()  {
	fmt.Printf("Hello #%v\n",counter)
	m.RUnlock()
	wg.Done()
}

func increment()  {
	counter++
	m.Unlock()	
	wg.Done()
}



