package configurator

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	ver = "v1"
)

// ConfigHandler is an http handler for a server
type ConfigHandler struct {
	S Server
}

// AddHandler wraps the AddTarget server request with an http handler
func (c *ConfigHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	module := mux.Vars(r)["module"]
	target := mux.Vars(r)["target"]

	// Find the corresponding module and append the target
	if err := c.S.AddTarget(target, module); err != nil {
		if err == ErrNotFound { // module wasn't found
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// DelHandler wraps the DelTarget server request with an http handler
func (c *ConfigHandler) DelHandler(w http.ResponseWriter, r *http.Request) {
	module := mux.Vars(r)["module"]
	target := mux.Vars(r)["target"]

	// Find the corresponding module and delete the target
	if err := c.S.DelTarget(target, module); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetHandler wraps the Get server request with an http handler
func (c *ConfigHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	out := c.S.Get()
	if err := json.NewEncoder(w).Encode(out); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetTargetsHandler returns all targets for a given module
func (c *ConfigHandler) GetTargetsHandler(w http.ResponseWriter, r *http.Request) {
	module := mux.Vars(r)["module"]
	s, err := c.S.GetTargets(module)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", s)
}
