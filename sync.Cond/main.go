package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	dishesReady = 0
)
// a cond  call is must be wrappend  inside a mutex lock and unlock

var (
	notificationReady = false
	notificationMessage string
	mu = sync.Mutex{}
	cond = sync.NewCond(&mu)
	wg = sync.WaitGroup{}
)

func waiter(id int) {
	defer wg.Done()
	mu.Lock()
	for dishesReady == 0 {
		fmt.Printf("Waiter %d is waiting for the dish to be ready\n", id)
		cond.Wait()
	}
	dishesReady--
	fmt.Printf("Waiter %d took the dish\n", id)
	mu.Unlock()
}

func cook(count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		mu.Lock()
		dishesReady++
		fmt.Printf("Cook prepared a dish %d\n", dishesReady)
		cond.Signal()
		mu.Unlock()
	}
}

func main() {
	waiterCount := 10
	chefCount := 10
	dishesPerChef := 1 


	for i := 0; i < waiterCount; i++ {
		wg.Add(1)
		go waiter(i)
	}

	for i := 0; i < chefCount; i++ {
		wg.Add(1)
		go cook(dishesPerChef)
	}

	wg.Wait()
	fmt.Println("All waiters have taken the dish")
	notification()
	
}

func client(id int) {
	defer wg.Done()
	mu.Lock()
	for !notificationReady {
		fmt.Printf("Client %d is waiting for the notification\n", id)
		cond.Wait()
	}
	fmt.Printf("Client %d received the notification: %s\n", id, notificationMessage)
	mu.Unlock()
}

func server() {	
	time.Sleep(time.Second * 2)
	mu.Lock()
	notificationReady = true
	notificationMessage = "New notification"
	cond.Broadcast()
	mu.Unlock()
}
func notification() {
	numClients := 10
	wg.Add(numClients)

	for i := 0; i < numClients; i++ {
		go client(i)
	}

	go server()

	wg.Wait()
	fmt.Println("All clients have received the notification")
}