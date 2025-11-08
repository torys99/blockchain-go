package main

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
import (
	"fmt"
	"sync"
	"time"
)

type Task func()

func scheduler(tasks []Task) {
	wg := sync.WaitGroup{}
	for i, task := range tasks {
		wg.Add(1)
		go func(task Task, index int) {
			defer wg.Done()
			start := time.Now()
			task()
			elapsed := time.Since(start)
			fmt.Printf("任务 %d 执行时间: %s\n", index, elapsed)
		}(task, i)
	}
	wg.Wait()
}

func main() {
	tasks := []Task{
		func() { time.Sleep(2 * time.Second) },
		func() { time.Sleep(1 * time.Second) },
		func() { time.Sleep(3 * time.Second) },
	}
	scheduler(tasks)
}
