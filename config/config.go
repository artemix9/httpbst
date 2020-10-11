package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Values []int
}

func Load(filename string) (Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("config: load: read file: %s", err)
	}

	conf := Config{}
	if err := json.Unmarshal(data, &conf); err != nil {
		return Config{}, fmt.Errorf("config: load: unmarshal: %s", err)
	}

	return conf, nil
}

func (conf Config) Validate() error {
	if conf.Values == nil {
		return fmt.Errorf("config: validate: values must be provided")
	}

	if len(conf.Values) != 30 {
		return fmt.Errorf("config: validate: values length must be 30")
	}

	return nil
}
