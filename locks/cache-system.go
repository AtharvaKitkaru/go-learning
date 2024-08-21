package locks

import (
    "fmt"
    "sync"
)

// Global WaitGroup for synchronization
var wg sync.WaitGroup

// Cache datatype
type Cache struct {
    data        string
    initLock    sync.Mutex // Lock used during initialization
    rwLock      sync.RWMutex // Lock for reading and writing
    cond        *sync.Cond // Condition variable for signaling
    initialized bool // Indicating if cache is initialized or not
}

// Init function
func (c *Cache) initCache(data string) {
    c.initLock.Lock()
    

    if !c.initialized {
        fmt.Println("Initializing Cache...")
        c.data = data
        c.initialized = true
        c.cond.Broadcast() // Notify all waiting readers
        fmt.Println("Cache Initialized.")
    }
	c.initLock.Unlock()
}

// Reader function
func (c *Cache) readerFunction() {
    defer wg.Done()

	c.initLock.Lock()
    for !c.initialized {
        c.cond.Wait() // Wait for the condition to be signaled
    }
	c.initLock.Unlock()

    c.rwLock.RLock()
    fmt.Println("Value of data is ", c.data)
	c.rwLock.RUnlock()
}

// Writer function
func (c *Cache) writerFunction(data string) {
    defer wg.Done()

    c.rwLock.Lock()
    c.initCache(data)
	c.rwLock.Unlock()
}

// Function that reads and writes to Cache
func SimulateCache() {
    cache := Cache{}
    cache.cond = sync.NewCond(&cache.initLock) // Use initLock for the condition variable

    // Create 4 reader goroutines
    for i := 0; i < 4; i++ {
        wg.Add(1)
        go cache.readerFunction()
    }

    // Create \ writer goroutines
  
        wg.Add(1)
        go cache.writerFunction(fmt.Sprintf("data %d", 1))


    wg.Wait()

}
