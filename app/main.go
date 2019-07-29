package main

import (
	"flag"
	"fmt"
	"log"

	proxy "../proxy"
)

var port = flag.Int("port", 8080, "the port the server will listen on")
var cpanelHostname = flag.String("cpanel-hostname", "", "the hostname of the cpanel")

func main() {
	flag.Parse()

	addr := fmt.Sprintf(":%v", *port)

	log.Fatal(proxy.Listen(addr, *cpanelHostname))
}
