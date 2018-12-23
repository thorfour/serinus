package main

import (
	"flag"
	"log"

	"github.com/thorfour/serinus/pkg/configurator"
)

var (
	cfgpath = flag.String("c", "/etc/prom.yml", "prometheus config file")
	port    = flag.Int("p", 9091, "port to server config sever on")
)

func init() {
	flag.Parse()
}

func main() {
	c, err := configurator.NewConfigServer(*cfgpath)
	if err != nil {
		log.Fatal("failed to load config server")
	}

	configurator.StartHTTPServer(c, *port)
}
