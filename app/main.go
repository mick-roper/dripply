package main

import (
	"flag"
	"fmt"
	"log"

	proxy "../proxy"
	targets "../proxy/targets"
)

const (
	blockSize = 32*1024
	bufferSize = blockSize * 2000 // 10,000 blocks
)

var port = flag.Int("port", 8080, "the port the server will listen on")
var cpanelHostname = flag.String("cpanel-hostname", "", "the hostname of the cpanel")

func main() {
	flag.Parse()

	addr := fmt.Sprintf(":%v", *port)

	targetCollection := targets.NewTargetCollection()
	targetCollection.SetTarget("localhost:8080", &targets.SimpleTarget{Hostname: "symmetric.solutions", Scheme: "https"})

	memoryBuffer, err := proxy.NewMemoryBuffer(bufferSize, blockSize)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(proxy.Listen(addr, *cpanelHostname, targetCollection, memoryBuffer))
}
