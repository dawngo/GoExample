// 官方方案
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			fmt.Println(ctx.Err())
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)

	time.Sleep(3 * time.Second)
	cancel() // 通知子goroutine退出

	wg.Wait()
	fmt.Println("over")
}
