# Concurrent Stock Exchange Simulator

This project simulates a simple stock exchange using Go's low-level concurrency primitives. It demonstrates how to handle concurrent reads and writes with `sync.RWMutex` to ensure thread-safe access to shared data structures.

## Features

- **Concurrent Reads:** Multiple goroutines can read the order book simultaneously.
- **Exclusive Writes:** Only one goroutine can modify (write to) the order book at a time.
- **Order Placement:** Supports placing buy and sell orders concurrently.
- **Order Cancellation:** Allows canceling orders by their ID.
- **Thread Safety:** All operations on the `OrderBook` are protected by a read-write mutex for safe concurrent access.

## Code Structure

- `Order`: Represents a buy or sell order with ID, price, amount, and side.
- `OrderBook`: Maintains slices of buy and sell orders and synchronizes access with `sync.RWMutex`.
- `PlaceOrder`: Adds a new order to the order book.
- `CancelOrder`: Cancels an order by ID (if present).
- `GetOrderBook`: Returns a snapshot of the current order book.
- `main`: Demonstrates concurrent order placement, reading, and cancellation using goroutines and a wait group.

## How It Works

1. Multiple goroutines place buy and sell orders concurrently.
2. The `OrderBook` uses an RWMutex:
   - **Lock** for writing (placing or canceling orders).
   - **RLock** for reading (retrieving the order book).
3. After all orders are placed, several goroutines read the current order book concurrently.
4. Some orders are canceled concurrently.
5. The final state of the order book is printed.

## Example Output

```
Placed buy order: {id:1 price:100 amount:10 side:buy}
Placed buy order: {id:2 price:101 amount:11 side:buy}
Placed sell order: {id:6 price:105 amount:15 side:sell}
...
Canceled buy order ID: 2
Cancel order 2 success: true
Canceled sell order ID: 7
Cancel order 7 success: true
[Reader 0] Buys: [{id:1 price:100 amount:10 side:buy} ...]
[Reader 0] Sells: [{id:6 price:105 amount:15 side:sell} ...]
Final Order Book:
Buys: [{id:1 price:100 amount:10 side:buy} ...]
Sells: [{id:6 price:105 amount:15 side:sell} ...]
```

## How to Run

1. Ensure you have Go installed (version 1.13+ recommended).
2. Clone this repository.
3. Navigate to the `stock-exchange` directory.
4. Run:

   ```sh
   go run main.go
   ```

## Concurrency Model

- The example demonstrates how to use Go's goroutines for concurrent operations and `sync.WaitGroup` to synchronize their completion.
- The use of `sync.RWMutex` allows for efficient, safe access patterns when reads are frequent and writes are less common.

## Notes

- The current implementation is a simplified simulation and does not include order matching or advanced trading logic.
- Order IDs are generated based on the current length of the order book and are not guaranteed to be globally unique in a distributed setting.

## License

This project is licensed under the MIT License.