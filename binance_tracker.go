package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type BinanceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func fetchPrice(symbol string, wg *sync.WaitGroup) {
	defer wg.Done()
	
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)
	
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("❌ 故障: 无法连接 Binance (%s)\n", symbol)
		return
	}
	defer resp.Body.Close()

	var result BinanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("❌ 解析失败: %s\n", symbol)
		return
	}

	fmt.Printf("🚀 锁定 | %s 实时价格: $%s\n", result.Symbol, result.Price)
}

func main() {
	targets := []string{
		"BTCUSDT", "ETHUSDT", "BNBUSDT", "SOLUSDT", "ADAUSDT",
		"XRPUSDT", "DOTUSDT", "DOGEUSDT", "AVAXUSDT", "LINKUSDT",
	}

	var wg sync.WaitGroup
	startTime := time.Now()

	fmt.Println("🛰️  发射并发抓取引擎...")

	for _, symbol := range targets {
		wg.Add(1)
		go fetchPrice(symbol, &wg)
	}

	wg.Wait()
	elapsed := time.Since(startTime)
	fmt.Printf("\n🌌 任务完成！总耗时: %s\n", elapsed)
}
