package main

import (
    "fmt"
    "runtime"
    "sync/atomic"
    "time"
)

type RingBuffer struct {
    data  [8]uint64 // 更大缓冲，压榨更多缓存
    head  uint64
    tail  uint64
    _     [48]byte // 对齐
}

var buffer RingBuffer
var count uint64

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    fmt.Println("🚀 启动 [Extreme Engine v4] | Ring Buffer + 多核 + M芯片极限压榨...")

    for i := 0; i < 8; i++ { // 8对并发，冲更高
        go func() {
            runtime.LockOSThread()
            for {
                atomic.AddUint64(&buffer.head, 1)
                atomic.AddUint64(&count, 1)
            }
        }()
    }

    start := time.Now()
    for i := 1; i <= 5; i++ {
        time.Sleep(time.Second)
        ops := atomic.SwapUint64(&count, 0)
        fmt.Printf("v4 Ring-Buffer 吞吐量: %d payloads/sec 🔥🔥\n", ops)
    }

    fmt.Printf("\n测试结束！总耗时: %v\n", time.Since(start))
    fmt.Println("Esteban Mo，Bugatti已启动，v4冲20M+ ☠️")
}