package dns

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type dnsityConfig struct {
	Records []Record `yaml:"records"`
}

type Record struct {
	Name string `yaml:"name"`
	Ip   string `yaml:"ip"`
}

func configLoader(filepath string) error {

	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		logger.Errorf("cannot read file : %s : %v", filepath, err)
	}

	var config dnsityConfig
	if err := yaml.Unmarshal(b, &config); err != nil {
		logger.Errorf("cannot unmarshal config : %v", err)
	}

	for _, r := range config.Records {
		RegisterRecord(r.Name, r.Ip)
	}

	return nil
}
