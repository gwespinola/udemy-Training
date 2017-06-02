package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter)) // access without race
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go

