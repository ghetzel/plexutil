package plexutil

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"strings"
)

type Configuration struct {
	URL        string                 `json:"url"`
	Parameters map[string]interface{} `json:"parameters"`
	Headers    map[string]interface{} `json:"headers"`
	path       string
}

func (self *Configuration) Save() error {
	if self.path != `` {
		if data, err := yaml.Marshal(self); err == nil {
			return ioutil.WriteFile(self.path, data, 0600)
		} else {
			return err
		}
	} else {
		return fmt.Errorf("Cannot save configuration without a given path")
	}
}

func DefaultConfig() Configuration {
	return Configuration{}
}

func LoadConfig(path string) (Configuration, error) {
	config := DefaultConfig()

	if strings.Contains(path, `~`) {
		if v := os.Getenv(`HOME`); v != `` {
			path = strings.Replace(path, `~`, v, 1)
		}
	}

	if data, err := ioutil.ReadFile(path); err == nil {
		config.path = path
		err := yaml.Unmarshal(data, &config)
		return config, err
	} else {
		return config, err
	}
}
