package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// PaddedCounter 结构体：通过 64 字节填充解决 False Sharing（伪共享）问题
// 这是顶级系统性能优化的“黑魔法”
type PaddedCounter struct {
	value uint64
	_     [56]byte // 填充 56 字节，加上 8 字节的 uint64，正好是一个 CPU 缓存行
}

func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	// 分配对齐后的计数器数组
	counters := make([]PaddedCounter, cores)

	fmt.Printf("🚀 硬件对齐引擎点火 | 锁定核心: %d | 准备压榨 M 芯片物理极限...\n", cores)

	for i := 0; i < cores; i++ {
		go func(id int) {
			runtime.LockOSThread() // 将 Goroutine 绑定到特定的物理线程
			for {
				// 原子操作，且每个核心只改自己的缓存行
				atomic.AddUint64(&counters[id].value, 1)
			}
		}(i)
	}

	start := time.Now()
	for {
		time.Sleep(time.Second)
		var totalOps uint64
		for i := 0; i < cores; i++ {
			// 读取并重置，计算每秒总吞吐量
			totalOps += atomic.SwapUint64(&counters[i].value, 0)
		}

		fmt.Printf("⚡️ 每秒吞吐量 (Throughput): %d ops/sec\n", totalOps)

		if time.Since(start) > 10*time.Second {
			fmt.Println("\n🏁 极限压力测试结束。这，就是代码优化的力量。")
			break
		}
	}
}
