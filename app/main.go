package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	proxy "../proxy"
	targets "../proxy/targets"
)

const (
	blockSize  = 32 * 1024
	bufferSize = blockSize * 2000 // 10,000 blocks
)

var port = flag.Int("port", 8080, "the port the server will listen on")
var cpanelHostname = flag.String("cpanel-hostname", "", "the hostname of the cpanel")

func main() {
	flag.Parse()

	addr := fmt.Sprintf(":%v", *port)

	targetCollection := targets.NewTargetCollection()
	targetCollection.SetTarget("localhost:8080", &targets.SimpleTarget{Hostname: "symmetric.solutions", Scheme: "https"})

	pool, err := proxy.NewPool(blockSize, 1000)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			printMemUsage()

			time.Sleep(1 * time.Second)
		}
	}()

	log.Fatal(proxy.Listen(addr, *cpanelHostname, targetCollection, pool))
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
