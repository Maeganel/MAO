package mud

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Address string
	Login   struct {
		Name     string
		Password string
	}
	Prompt    Pattern
	Abilities map[string]struct {
		Ready []string
		Wait  []string
	}
	Triggers map[Pattern]string
	Vars     map[string]string
	Aliases  map[string]string
	Log      map[string]struct {
		Timestamp bool               `yaml:"timestamp,omitempty"`
		Match     map[Pattern]string `yaml:"match,omitempty"`
	} `yaml:"log,omitempty"`
	Dump map[string]struct {
		Cmd   string
		Dest  string
		Start Pattern
		End   Pattern
	}
	Highlight map[Pattern]*Color
	Replace   map[Pattern]struct {
		With  string
		Color *Color
	}
	Gag    []Pattern
	Timers []struct {
		Every time.Duration
		Do    string
	}
	Tick struct {
		Duration time.Duration
		Match    []string
	}
}

func UnmarshalConfig(path string) (Config, error) {
	var cfg Config

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(buf, &cfg)
	return cfg, err
}
