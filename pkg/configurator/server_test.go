package configurator

import (
	"net/url"
	"os"
	"sync"
	"testing"

	"github.com/prometheus/prometheus/config"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Remove("http_2xx.json")
	os.Exit(code)
}

func TestAddTarget(t *testing.T) {
	c := new(ConfigServer)
	c.filelocks = make(map[string]*sync.Mutex)
	c.filelocks["http_2xx.json"] = new(sync.Mutex)

	c.Config = &config.Config{
		ScrapeConfigs: []*config.ScrapeConfig{
			&config.ScrapeConfig{
				JobName: "test",
				Params: url.Values{
					"module": []string{"http_2xx"},
				},
			},
		},
	}

	require.NoError(t, c.AddTarget("digitalocean.com", "http_2xx"))

	targets, err := c.GetTargets("http_2xx")
	require.NoError(t, err)

	require.Equal(t, []string{"digitalocean.com"}, targets)
}

func TestDelTarget(t *testing.T) {
	c := new(ConfigServer)
	c.filelocks = make(map[string]*sync.Mutex)
	c.filelocks["http_2xx.json"] = new(sync.Mutex)

	c.Config = &config.Config{
		ScrapeConfigs: []*config.ScrapeConfig{
			&config.ScrapeConfig{
				JobName: "test",
				Params: url.Values{
					"module": []string{"http_2xx"},
				},
			},
		},
	}

	require.NoError(t, c.AddTarget("digitalocean.com", "http_2xx"))
	targets, err := c.GetTargets("http_2xx")
	require.NoError(t, err)
	require.Equal(t, []string{"digitalocean.com"}, targets)
	require.NoError(t, c.DelTarget("digitalocean.com", "http_2xx"))
	targets, err = c.GetTargets("http_2xx")
	require.NoError(t, err)

	require.Equal(t, []string(nil), targets)
}

func TestLoad(t *testing.T) {
	c, err := NewConfigServer("prom.yml")
	require.NoError(t, err)
	require.Equal(t, c.Config.ScrapeConfigs[0].JobName, "blackbox")
}
