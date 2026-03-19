package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// 🛡️ 硬件对齐：强行填充 64 字节，彻底杜绝“伪共享”
type ControlSignal struct {
	flag uint32
	_    [60]byte 
}

type DataPayload struct {
	ID   uint64
	Data [8]uint64 
	_    [56]byte 
}

func main() {
	// 锁定两个核心进行极限肉搏
	runtime.GOMAXPROCS(2)

	var signal ControlSignal
	var cargo DataPayload
	var count uint64

	fmt.Println("🛰️  启动 [缓存隔离] 传输引擎 | 正在压榨 M 芯片物理极限...")

	// 生产者：独占一个核心，疯狂写入
	go func() {
		runtime.LockOSThread()
		for {
			for atomic.LoadUint32(&signal.flag) != 0 {}
			cargo.ID++ 
			atomic.StoreUint32(&signal.flag, 1) 
		}
	}()

	// 消费者：独占另一个核心，疯狂读取
	go func() {
		runtime.LockOSThread()
		for {
			for atomic.LoadUint32(&signal.flag) != 1 {}
			_ = cargo.ID 
			atomic.AddUint64(&count, 1)
			atomic.StoreUint32(&signal.flag, 0) 
		}
	}()

	start := time.Now()
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		ops := atomic.SwapUint64(&count, 0)
		fmt.Printf("🚀 隔离后吞吐量: %d payloads/sec\n", ops)
	}

	duration := time.Since(start)
	fmt.Printf("\n🏁 测试结束。总耗时: %v\n", duration)
	fmt.Println("Esteban，如果数字大幅提升，恭喜你，你已经掌握了‘欺骗硬件’的顶级艺术。")
}
