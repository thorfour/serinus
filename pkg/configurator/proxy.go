package configurator

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Region is a router for routing configurations based on region
type Region map[string]string

// Director is a httputil.ReverseProxy director function.
// It strips the region query param and replaces the host with the regions specific host
func (r Region) Director(req *http.Request) {
	region := mux.Vars(req)["region"]
	q := req.URL.Query()
	q.Del("region")

	req.URL.RawQuery = q.Encode()
	req.URL.Scheme = "http"
	req.URL.Host = r[region]
}
