package configurator

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

// StartHTTPServer starts an HTTP configurator server
func StartHTTPServer(c Server, port int) {
	h := ConfigHandler{S: c}
	r := mux.NewRouter()

	subRouter := r.PathPrefix("/api/v1").Subrouter()
	subRouter.Methods("PATCH").Path("/add").Queries("target", "{target}", "module", "{module}").Handler(http.HandlerFunc(h.AddHandler))
	subRouter.Methods("DELETE").Path("/del").Queries("target", "{target}", "module", "{module}").Handler(http.HandlerFunc(h.DelHandler))
	subRouter.Methods("GET").Path("/get").Handler(http.HandlerFunc(h.GetHandler))
	subRouter.Methods("GET").Path("/targets").Queries("module", "{module}").Handler(http.HandlerFunc(h.GetTargetsHandler))

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
		log.Fatal("config server exited with error")
	}
}

// StartHTTPProxy starts a configurator proxy
func StartHTTPProxy(p *httputil.ReverseProxy, port int) {
	r := mux.NewRouter()

	subRouter := r.PathPrefix("/api/v1").Subrouter()
	subRouter.Methods("PATCH").Path("/add").Queries("target", "{target}", "module", "{module}", "region", "{region}").Handler(http.HandlerFunc(p.ServeHTTP))
	subRouter.Methods("DELETE").Path("/del").Queries("target", "{target}", "module", "{module}", "region", "{region}").Handler(http.HandlerFunc(p.ServeHTTP))
	subRouter.Methods("GET").Path("/get").Queries("region", "{region}").Handler(http.HandlerFunc(p.ServeHTTP))
	subRouter.Methods("GET").Path("/targets").Queries("module", "{module}", "region", "{region}").Handler(http.HandlerFunc(p.ServeHTTP))

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
		log.Fatal("config server exited with error")
	}
}
