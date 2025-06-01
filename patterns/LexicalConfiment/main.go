package main


import (
	"fmt"
)

//Lexical Confiment is a pattern that ensures that the order of execution of goroutines is deterministic.


func main() {
	// chanOwner creates and owns the write aspect of the channel.
	// It confines write access to within its lexical scope.
	chanOwner := func() <-chan int {
		// This channel is created within the chanOwner function's lexical scope.
		// Only this function and its closure have write access to it.
		results := make(chan int, 5)

		// A goroutine is launched that exclusively writes to the channel.
		go func() {
			defer close(results) // The channel is closed after all writes are done.
			for i := 1; i <= 5; i++ {
				results <- i // Write access is confined to this goroutine only.
			}
		}()

		// Only the read-only aspect of the channel is returned,
		// preventing any external goroutine from writing to it.
		return results
	}

	// Here, the main goroutine receives only the read-only view of the channel.
	result := chanOwner()

	// The main goroutine can only read from the channel, not write to it.
	for result := range result {
		fmt.Println(result)
	}
	fmt.Println("Done")
}


// | Feature / Case                      | Lexical Confinement ✅ | Synchronization 🛠        | Channels 📡                |
// | ----------------------------------- | --------------------- | ------------------------- | -------------------------- |
// | No shared state                     | ✅ Ideal               | ❌ Unnecessary             | ✅ Possible, but overkill   |
// | Split data and assign per goroutine | ✅ Perfect fit         | ❌ Overhead                | ✅ Optional                 |
// | Shared map / counter                | ❌ Not possible        | ✅ Use Mutex/Atomic        | ✅ If modeled correctly     |
// | Wait for multiple goroutines        | ❌                     | ✅ sync.WaitGroup          | ✅ Can signal via channel   |
// | Avoid race conditions               | ✅ Naturally race-free | ⚠️ Requires discipline    | ✅ Naturally safe           |
// | Performance (no locks)              | ✅                     | ⚠️ May degrade with locks | ✅ Efficient for small data |
// | Coordination between goroutines     | ❌                     | ❌                         | ✅ Natural choice           |
