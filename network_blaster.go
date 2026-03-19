package main

import (
	"fmt"
	"net"
	"runtime"
)

func main() {
	// 锁定全核心，进入狂暴发射模式
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 目标：你刚才那个监听 9000 端口的“接收站”
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:9000")
	
	fmt.Println("🚀 [网络轰炸机] 修复成功！正在全速倾泄数据包...")

	// 启动与核心数相等的发射引擎
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			conn, _ := net.DialUDP("udp", nil, addr)
			payload := []byte("ESTEBAN_POWER") // 带着你的名字发射
			for {
				// 没有任何 Sleep，死循环压榨总线
				conn.Write(payload) 
			}
		}()
	}

	// 保持运行，直到你按下 Control + C
	select {}
}
