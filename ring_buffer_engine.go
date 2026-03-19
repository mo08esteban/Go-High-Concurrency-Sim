package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// 模拟一个巨大的“货舱”
type DataPayload struct {
	ID        uint64
	Timestamp int64
	Data      [8]uint64 // 模拟 64 字节的硬核数据
}

func main() {
	runtime.GOMAXPROCS(2)

	// 建立一个物理上的“共享货舱”
	var cargo DataPayload
	var signal uint32 // 0: 生产中, 1: 已装载
	var count uint64

	fmt.Println("🚀 启动 [零拷贝] 环形传输引擎...")

	// 生产者：直接在物理内存上改写，不复制
	go func() {
		runtime.LockOSThread()
		for {
			if atomic.LoadUint32(&signal) == 0 {
				cargo.ID++ // 直接修改物理内存
				cargo.Timestamp = time.Now().UnixNano()
				atomic.StoreUint32(&signal, 1) // 发射信号
			}
		}
	}()

	// 消费者：直接从同一块物理内存读取
	go func() {
		runtime.LockOSThread()
		for {
			if atomic.LoadUint32(&signal) == 1 {
				_ = cargo.ID // 直接读取，零复制
				atomic.AddUint64(&count, 1)
				atomic.StoreUint32(&signal, 0) // 重置信号
			}
		}
	}()

	start := time.Now()
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		ops := atomic.SwapUint64(&count, 0)
		fmt.Printf("📦 零拷贝吞吐量: %d payloads/sec\n", ops)
	}

	fmt.Printf("\n🏁 极限测试结束！总耗时: %v\n", time.Since(start))
	fmt.Println("如果你能优化到 3000 万以上，你就是顶级的系统架构师。")
}
