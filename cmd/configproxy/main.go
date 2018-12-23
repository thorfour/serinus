package main

import (
	"flag"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/thorfour/serinus/pkg/configurator"
)

var (
	port      = flag.Int("p", 9091, "port to serve config router on")
	endpoints = flag.String("r", "", "comman separated list of region endpoint string pairs (ex: 'nyc3, http://mynyc3endpint')")
)

func init() {
	flag.Parse()
}

func main() {
	r := buildRegionTable(*endpoints)
	proxy := &httputil.ReverseProxy{
		Director: r.Director,
	}

	configurator.StartHTTPProxy(proxy, *port)
}

func buildRegionTable(list string) configurator.Region {
	pairs := strings.Split(list, ",")
	if len(pairs)%2 != 0 || len(pairs) == 0 {
		logrus.Error("invalid number of regions require an even list")
		os.Exit(1)
	}

	// Build region lookup table
	region := configurator.Region{}
	for i, s := range pairs {
		if i%2 == 1 {
			region[pairs[i-1]] = s
		}
	}

	return region
}
