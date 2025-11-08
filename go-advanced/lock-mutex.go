/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	numGoroutines := 10
	incrementsPerGoroutine := 1000
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter Value:", counter)
}
