//Graceful handling of anamolies in go routines


package main

	import (
	"fmt"
	"time"
)





func main() {
	doWork := func(done <-chan bool, strings <-chan string) <-chan string {
		dataStream := make(chan string)

		go func() {
			defer fmt.Println("doWork exited")
			defer close(dataStream)

			for {
				select {
				case s, ok := <-strings:
					if !ok {
						return
					}
					dataStream <- s
				case <-done:
					return
				}
			}
		}()

		return dataStream
	}

	done := make(chan bool)
	strings := make(chan string)


	dataStream := doWork(done, strings)


	go func() {
		for i := 0; i < 10; i++ {
			strings <- "Hello, World!"
		}
		close(strings)
	}()


	go func() {
		for d := range dataStream {
			fmt.Println("Received:", d)
		}
	}()


	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
