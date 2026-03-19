package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	// 锁定 2 个核心进行“肉搏”测试
	runtime.GOMAXPROCS(2)

	var ringBuffer uint64 
	var count uint64
	
	fmt.Println("🚀 启动双核心 [无锁传输] 引擎...")

	// 核心 1：生产者
	go func() {
		runtime.LockOSThread()
		for {
			if atomic.CompareAndSwapUint64(&ringBuffer, 0, 1) {}
		}
	}()

	// 核心 2：消费者
	go func() {
		runtime.LockOSThread()
		for {
			if atomic.CompareAndSwapUint64(&ringBuffer, 1, 0) {
				atomic.AddUint64(&count, 1) 
			}
		}
	}()

	// 记录起始时间
	start := time.Now()

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		ops := atomic.SwapUint64(&count, 0)
		fmt.Printf("📦 第 %d 秒 | 核心间吞吐量: %d packets/sec\n", i, ops)
	}
	
	// 使用 start 变量计算总耗时，让编译器闭嘴并展示结果
	duration := time.Since(start)
	fmt.Printf("\n🏁 极限测试结束！实际运行时间: %v\n", duration)
	fmt.Println("这是工业级无锁通信的真实速度。")
}
