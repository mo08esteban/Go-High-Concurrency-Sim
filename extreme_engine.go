package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

type DataPayload struct {
	ID   uint64
	Data [8]uint64 
}

func main() {
	// 锁定 2 个核心，我们要看纯粹的电子交换速度
	runtime.GOMAXPROCS(2)

	var cargo DataPayload
	var signal uint32 
	var count uint64

	fmt.Println("⚡️ 启动 [极致物理] 传输引擎 | 移除所有系统调用...")

	// 生产者：纯内存操作
	go func() {
		runtime.LockOSThread()
		for {
			// 自旋等待，直到信号为 0
			for atomic.LoadUint32(&signal) != 0 {}
			
			cargo.ID++ 
			// 这里不再调用 time.Now()，只做纯粹的内存搬运
			atomic.StoreUint32(&signal, 1) 
		}
	}()

	// 消费者：纯内存操作
	go func() {
		runtime.LockOSThread()
		for {
			// 自旋等待，直到信号为 1
			for atomic.LoadUint32(&signal) != 1 {}
			
			_ = cargo.ID 
			atomic.AddUint64(&count, 1)
			atomic.StoreUint32(&signal, 0) 
		}
	}()

	start := time.Now()
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		ops := atomic.SwapUint64(&count, 0)
		fmt.Printf("🔥 物理吞吐量: %d payloads/sec\n", ops)
	}

	fmt.Printf("\n🏁 极限测试结束。总耗时: %v\n", time.Since(start))
	fmt.Println("这才是马斯克员工讨论的‘无损吞吐’。")
}
