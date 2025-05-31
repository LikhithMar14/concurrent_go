package main

import (
	"fmt"
)

/*
	Syntax of channels

	var channel chan int //default is nil not initialized

	channel = make(chan interface{}) //initialized

	//Right side => Send only
	//Left side => Receive only

	only receive channel
	channel  =  make(<- chan interface{})

	only send channel

	channel = make(chan <- interface{})


	Go will implicitly convert bidirectional channels to unidirectional
	channels when needed. Here’s an example:


	var receiveChan <-chan interface{}
	var sendChan chan<- interface{}
	dataStream := make(chan interface{})

	receiveChan = dataStream
	sendChan = dataStream


*/

func main() {
	// stringStream := make(chan string)

	// go func(){
	// 	stringStream <- "Hello channels!"
	// }()

	// so ok if send channel hits first it waits until and unless some one receives it , and if receive channel hits first it waits until and unless some one sends to it

	/*
		stringStream := make(chan string)
			go func() {
			if 0 != 1 {
			return
			}
			stringStream <- "Hello channels!"
			}()
			fmt.Println(<-stringStream)

			when you are done with writing to a channel we should closeit using close(channel)

			then in ok we will get false this is the correct way to close the channels and inform down stream reciever that the channel is closed


	*/

	// salutation, ok := <-stringStream
	// fmt.Println(salutation, ok)

	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for i := range 10 {
			intStream <- i
		}
	}()

	for i := range intStream {
		fmt.Println(i)
	}

	fmt.Println("Done")

	/*
		In our main goroutine, we prepare some resources or do setup work first. After that preparation is complete, we close the channel to signal all waiting goroutines to proceed with their work.

		then in the for loop we will get the value from the channel and print itpackage main

					import (
							"fmt"
							"sync"
					)

					func main(){
						dataStream := make(chan int)
						var wg sync.WaitGroup

						for i:= 0; i <5; i++ {
							wg.Add(1)
							go func(i int){
								defer wg.Done()
								<-dataStream
								fmt.Println("Unblocking dataStream of goroutine: ",i)
							}(i)
						}
						fmt.Println("Unblocking goroutines...")
						close(dataStream)
						wg.Wait()
					}




		// If we use low level golang primitives we can achieve the same thing

		func main() {
			var mu sync.Mutex
			cond := sync.NewCond(&mu)
			var wg sync.WaitGroup

			const goroutines = 3

			for i := 0; i < goroutines; i++ {
				wg.Add(1)
				go func(id int) {
					defer wg.Done()

					cond.L.Lock()
					cond.Wait() // wait for signal
					cond.L.Unlock()

					fmt.Println("Goroutine", id, "started")
				}(i)
			}

			// Give goroutines time to start and wait
			fmt.Println("All goroutines waiting...")

			// Broadcast to wake all waiting goroutines
			cond.Broadcast()

			wg.Wait()
			}


			var stdoutBuff bytes.Buffer
			defer stdoutBuff.WriteTo(os.Stdout) // Write buffer content to stdout when done

			intStream := make(chan int, 4) // Buffered channel with capacity 4

			go func() {
			defer close(intStream)
			defer fmt.Fprintln(&stdoutBuff, "Producer Done.")

			for i := 0; i < 5; i++ {
				fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
				intStream <- i
			}
			}()

			for integer := range intStream {
			fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
			}

		// Using bytes.Buffer collects output from multiple goroutines without mixing it up.
		// Printing all at once keeps output ordered and clear. Direct fmt.Printf can jumble output when goroutines print concurrently.
	*/
}
