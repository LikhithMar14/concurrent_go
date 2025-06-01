// Simulating a stock exchange with golang's low level primitives where there can concurrnet reads but only one write at a time

// we use sync.RWMutex to simulate the stock exchange




package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	id     int
	price  float64
	amount int
	side   string
}

type OrderBook struct {
	mu    sync.RWMutex
	buys  []Order
	sells []Order
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		buys:  make([]Order, 0),
		sells: make([]Order, 0),
	}
}

func (ob *OrderBook) PlaceOrder(side string, price float64, amount int) int {
	//This lock is for writing to the order book so that no other goroutine can read or write to the order book concurrently
	ob.mu.Lock()
	defer ob.mu.Unlock()

	order := Order{
		id:     len(ob.buys) + len(ob.sells) + 1,
		price:  price,
		amount: amount,
		side:   side,
	}

	if side == "buy" {
		ob.buys = append(ob.buys, order)
	} else {
		ob.sells = append(ob.sells, order)
	}

	fmt.Printf("Placed %s order: %+v\n", side, order)
	return order.id
}

//user can only canel his own orders not others
func (ob *OrderBook) CancelOrder(id int) bool {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	for i, order := range ob.buys {
		if order.id == id {
			ob.buys = append(ob.buys[:i], ob.buys[i+1:]...)
			fmt.Printf("Canceled buy order ID: %d\n", id)
			return true
		}
	}

	for i, order := range ob.sells {
		if order.id == id {
			ob.sells = append(ob.sells[:i], ob.sells[i+1:]...)
			fmt.Printf("Canceled sell order ID: %d\n", id)
			return true
		}
	}

	return false
}

func (ob *OrderBook) GetOrderBook() (buys []Order, sells []Order) {
	ob.mu.RLock()
	defer ob.mu.RUnlock()

	
	buys = make([]Order, len(ob.buys))
	sells = make([]Order, len(ob.sells))
	copy(buys, ob.buys)
	copy(sells, ob.sells)
	return
}

func main() {
	ob := NewOrderBook()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ob.PlaceOrder("buy", 100+float64(i), 10+i)
		}(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ob.PlaceOrder("sell", 105+float64(i), 15+i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			buys, sells := ob.GetOrderBook()
			fmt.Printf("[Reader %d] Buys: %+v\n", i, buys)
			fmt.Printf("[Reader %d] Sells: %+v\n", i, sells)
		}(i)
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		success := ob.CancelOrder(2)
		fmt.Println("Cancel order 2 success:", success)
	}()
	go func() {
		defer wg.Done()
		success := ob.CancelOrder(7)
		fmt.Println("Cancel order 7 success:", success)
	}()

	wg.Wait()


	buys, sells := ob.GetOrderBook()
	fmt.Println("Final Order Book:")
	fmt.Printf("Buys: %+v\n", buys)
	fmt.Printf("Sells: %+v\n", sells)


	time.Sleep(time.Millisecond * 100)
}
