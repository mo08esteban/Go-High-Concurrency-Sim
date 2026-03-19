package main

import (
	"fmt"
	"net"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 开启一个 UDP 端口，像卫星接收站一样监听全球数据
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 9000,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		fmt.Printf("❌ 发射场故障: %v\n", err)
		return
	}
	defer conn.Close()

	var packetCount uint64
	fmt.Println("🚀 [Esteban Starship Net] 已上线！监听端口: 9000")
	fmt.Println("正在准备接收行星级数据流...")

	// 开启高并发接收引擎
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			buf := make([]byte, 1024)
			for {
				// 极致的阻塞接收
				_, _, err := conn.ReadFromUDP(buf)
				if err == nil {
					atomic.AddUint64(&packetCount, 1)
				}
			}
		}()
	}

	// 性能监控
	for {
		time.Sleep(time.Second)
		p := atomic.SwapUint64(&packetCount, 0)
		fmt.Printf("📡 当前网络吞吐: %d packets/sec (网络重力下，这个数字会很难看)\n", p)
	}
}
