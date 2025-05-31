package main

import (
	"fmt"
	"sync"
)

/*
	In Go, the garbage collector (GC) only reclaims memory that is no longer referenced. If you allocate large objects (like 1MB each) and store them in a slice, the GC won’t clean them up because the slice still holds references. This leads to memory not being freed even when you're logically done using it. To solve this, sync.Pool is used. It allows temporary object reuse without holding strong references. When you Put() an object into a pool, it becomes eligible for GC if not reused soon, reducing memory pressure. sync.Pool is ideal for managing short-lived, reusable large objects — it improves performance by minimizing allocations and letting GC clean up unused memory efficiently.
*/

/*
	memory

	sync.pool

	.GET
	.PUT

*/

type SampleObject struct {
	ID int
}

var idCounter int

func CreateSampleObject() *SampleObject {
	idCounter++
	return &SampleObject{ID: idCounter}
}

// use this in heavy operations
/*


 */
func main() {
	var memoryPiece int
	objectPool := sync.Pool{
		New: func() any {
			memoryPiece++
			return CreateSampleObject()
		},
	}
	const worker = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(worker)
	for range worker {
		go func() {
			obj := objectPool.Get().(*SampleObject) // with optimzation => 9 // without optimization = 1024*1024
			// Simulate using the object
			objectPool.Put(obj) // Immediately put back after use to handle memory leaks
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Completed Execution with sync.Pool ", memoryPiece)
}
