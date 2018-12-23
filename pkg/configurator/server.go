package configurator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/prometheus/prometheus/config"
)

var (
	// ErrNotFound is retured when the given module is not found
	ErrNotFound = fmt.Errorf("Module Not Found")
)

// Server is the interface for a configurator server
type Server interface {
	AddTarget(target, module string) error
	DelTarget(target, module string) error
	AddScrapeConfig(*config.ScrapeConfig) error
	Get() []*config.ScrapeConfig
	GetTargets(module string) ([]string, error)
}

// ConfigServer implements the Server interface
type ConfigServer struct {
	filelocks map[string]*sync.Mutex
	path      string
	Config    *config.Config
}

type targets struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

// NewConfigServer initializes a new config sever loading the config file from path
func NewConfigServer(path string) (*ConfigServer, error) {
	c := new(ConfigServer)
	c.path = path
	c.filelocks = make(map[string]*sync.Mutex)

	// Load configs from file
	b, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return c, nil
		}
		return nil, err
	}

	c.Config, err = config.Load(string(b))
	if err != nil {
		return nil, err
	}

	for _, cfg := range c.Config.ScrapeConfigs {
		c.filelocks[fmt.Sprintf("%s.json", cfg.Params.Get("module"))] = new(sync.Mutex)
	}
	return c, err
}

// AddTarget will add a target to the given module and write the file back out
func (c *ConfigServer) AddTarget(target, module string) error {

	// Find the corresponding module and append the target
	for _, cfg := range c.Config.ScrapeConfigs {
		if cfg.Params.Get("module") == module {
			return c.addTargetToFile(fmt.Sprintf("%s.json", module), target)
		}
	}

	return ErrNotFound
}

func (c *ConfigServer) addTargetToFile(filename, target string) error {
	_, err := c.readModifyFile(true, filename, func(list []*targets) []*targets {
		for _, g := range list { // Don't add a preexisting target
			for _, t := range g.Targets {
				if t == target {
					return list
				}
			}
		}

		if len(list) == 0 {
			list = append(list, &targets{
				Targets: []string{target},
			})
			return list
		}

		// NOTE: Only single group of targets supported
		list[0].Targets = append(list[0].Targets, target)
		return list
	})

	return err
}

func (c *ConfigServer) readModifyFile(saveChanges bool, filename string, f func([]*targets) []*targets) ([]*targets, error) {
	c.filelocks[filename].Lock()
	defer c.filelocks[filename].Unlock()

	b, err := ioutil.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	var list []*targets
	if !os.IsNotExist(err) {
		if err := json.Unmarshal(b, &list); err != nil {
			return nil, err
		}
	}

	if f != nil {
		list = f(list)
	}

	if saveChanges {
		b, err = json.Marshal(list)
		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(filename, b, 0666); err != nil {
			return nil, err
		}
	}

	return list, nil
}

// DelTarget removes a target from a
func (c *ConfigServer) DelTarget(target, module string) error {

	// Find the corresponding module and append the target
	for _, cfg := range c.Config.ScrapeConfigs {
		if cfg.Params.Get("module") == module {
			return c.delTargetFromFile(fmt.Sprintf("%s.json", module), target)
		}
	}

	return nil
}

func (c *ConfigServer) delTargetFromFile(filename, target string) error {
	_, err := c.readModifyFile(true, filename, func(list []*targets) []*targets {
		var i, j int
		var t string
		for i = range list {
			for j, t = range list[i].Targets {
				if t == target {
					break
				}
			}
		}

		list[i].Targets = append(list[i].Targets[:j], list[i].Targets[j+1:]...)

		return list
	})

	return err
}

// Get returns all current configs
func (c *ConfigServer) Get() []*config.ScrapeConfig {
	return c.Config.ScrapeConfigs
}

// GetTargets returns a list of targets for a given module
func (c *ConfigServer) GetTargets(module string) ([]string, error) {
	// Find the corresponding module and append the target
	for _, cfg := range c.Config.ScrapeConfigs {
		if cfg.Params.Get("module") == module {
			return c.getTargetsFromFile(fmt.Sprintf("%s.json", module))
		}
	}

	return nil, nil
}

func (c *ConfigServer) getTargetsFromFile(filename string) ([]string, error) {
	list, err := c.readModifyFile(false, filename, nil)
	if err != nil {
		return nil, err
	}

	var targets []string
	for i := range list {
		for _, a := range list[i].Targets {
			targets = append(targets, string(a))
		}
	}

	return targets, err
}

// AddScrapeConfig adds a new scrape config to the server
func (c *ConfigServer) AddScrapeConfig(sc *config.ScrapeConfig) error {
	return nil
}
