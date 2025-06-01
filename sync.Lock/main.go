package main

import (
    "fmt"
    "sync"
)

// func main() {
//     // var count int
//     // var lock sync.Mutex
//     // var wg sync.WaitGroup
    
//     // increment := func(){
//     //     defer wg.Done()
//     //     lock.Lock()
//     //     count++
//     //     lock.Unlock()
//     // }
//     // for range 100000 {
//     //     wg.Add(1)
        
//     //     go increment()
//     // }
//     // wg.Wait()
//     // fmt.Println(count)


	

// }

	


func main() {
    var count int
    var lock sync.Mutex
    
    increment := func() {
        lock.Lock()
        defer lock.Unlock()
        count++
        fmt.Printf("Incrementing: %d\n",count)
    }
    
    decrement := func() {
        lock.Lock()
        defer lock.Unlock()
        count--
        fmt.Printf("Decrementing: %d\n",count)
    }
    
    var arithmetic sync.WaitGroup
    for i := 0; i  <= 5; i++ {
        arithmetic.Add(1)
        go func() {
            defer arithmetic.Done()
            increment()
        }()
    }
    
    
    for i := 0; i <= 5; i++ {
        arithmetic.Add(1)
        go func() {
            defer arithmetic.Done()
            decrement()
        }()
    }
    
    
    arithmetic.Wait()
    
    fmt.Println(count)

	/*  
		sync.RWMutex is a kind of mutex where mulitple go routines can read the data concurrently but when a go routine want to write the data it will lock the mutex for the whole duration of the write operation

		so if we want to read the data concurrently we can use sync.RWMutex

		sync.RWMutex has two methods Lock and RLock

		Lock is used to lock the mutex for writing

		RLock is used to lock the mutex for reading

		so if we want to read the data concurrently we can use sync.RWMutex

		sync.RWMutex has two methods Lock and RLock

		Lock is used to lock the mutex for writing
	*/
}