package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// 模拟一笔交易订单
type Order struct {
	Price  uint64
	Amount uint64
}

func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	var matchedCount uint64
	var totalVolume uint64

	fmt.Printf("🏦 启动 [Esteban 极速交易所] 核心 | 核心数: %d\n", cores)
	fmt.Println("正在模拟全球高频买卖撮合...")

	// 开启多个撮合引擎核心
	for i := 0; i < cores; i++ {
		go func() {
			runtime.LockOSThread()
			for {
				// 模拟撮合逻辑：买入价 == 卖出价
				// 在真实场景，这里会有复杂的红黑树或跳表操作
				buyOrder := Order{Price: 70000, Amount: 1}
				sellOrder := Order{Price: 70000, Amount: 1}

				if buyOrder.Price == sellOrder.Price {
					// 撮合成功
					atomic.AddUint64(&matchedCount, 1)
					atomic.AddUint64(&totalVolume, buyOrder.Amount)
				}
			}
		}()
	}

	start := time.Now()
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		m := atomic.SwapUint64(&matchedCount, 0)
		v := atomic.SwapUint64(&totalVolume, 0)
		fmt.Printf("🔔 第 %d 秒 | 撮合成功: %d 笔 | 成交量: %d BTC\n", i, m, v)
	}

	fmt.Printf("\n🏁 压力测试完成。总耗时: %v\n", time.Since(start))
	fmt.Println("这就是支撑全球金融系统的‘数字心脏’。")
}
