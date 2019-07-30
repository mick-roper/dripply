package main

import (
	"flag"
	"fmt"
	"log"

	proxy "../proxy"
	targets "../proxy/targets"
)

var port = flag.Int("port", 8080, "the port the server will listen on")
var cpanelHostname = flag.String("cpanel-hostname", "", "the hostname of the cpanel")

func main() {
	flag.Parse()

	addr := fmt.Sprintf(":%v", *port)

	targetCollection := &targets.TargetCollection{}

	targetCollection.SetTarget("localhost:8080", &targets.SimpleTarget{Hostname: "symmetric.solutions", Scheme: "https"})

	log.Fatal(proxy.Listen(addr, *cpanelHostname, targetCollection))
}
