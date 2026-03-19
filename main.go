package main

import (
	"fmt"
	"sync"
	"time"
)

func launchEngine(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 任务完成时告诉系统
	fmt.Printf("Engine #%d: Ignition Sequence Start... 🚀\n", id)
	time.Sleep(1 * time.Second) // 模拟点火耗时
	fmt.Printf("Engine #%d: Thrust is Stable! ✅\n", id)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("Mission Control: Starting Main Engine Ignition...")

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go launchEngine(i, &wg) // 这个 "go" 关键字就是并发的魔法
	}

	wg.Wait() // 等待所有引擎点火完成
	fmt.Println("SpaceX-Project: ALL ENGINES GO! WE HAVE LIFT OFF! 🌌")
}
