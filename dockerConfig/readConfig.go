package dockerConfig

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type DockerConfig struct {
	Course string `yaml:"course"`
	OperatingSystem string `yaml:"operatingSystem"`
	ProgrammingLanguage string `yaml:"programmingLanguage"`
}

func GetConfig(configFilename string) (*DockerConfig, error) {
	yamlFile, err := ioutil.ReadFile(configFilename)
	if err != nil {
		//return nil, errors.Wrap(err, "failed to read log config file")
	}

	c := &DockerConfig{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Error in unmarshalling yaml: %v", err)
	}
	return c, nil
}