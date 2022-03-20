package args

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var Config = &config{}
var builderConf = &configBuilder{config: Config}

// Argument config structure. It is private to make sure that only 1 instance can be created. It holds all
// configurations values passed to Dashboard binary.
type config struct {
	Database struct {
		Host     string `yaml:"dbHost"`
		Username string `yaml:"dbUsername"`
		Password string `yaml:"dbPassword"`
		Port     string `yaml:"dbPort"`
	} `yaml:"database"`
}

type configBuilder struct {
	config *config
}

// GetHostDatabase database host configurations of Dashboard binary.
func (self *config) GetHostDatabase() string {
	return self.Database.Host
}

// GetUsernameDatabase database username configurations of Dashboard binary.
func (self *config) GetUsernameDatabase() string {
	return self.Database.Username
}

// GetPasswordDatabase database password configurations of Dashboard binary.
func (self *config) GetPasswordDatabase() string {
	return self.Database.Password
}

// GetPortDatabase database port configurations of Dashboard binary.
func (self *config) GetPortDatabase() string {
	return self.Database.Port
}

// SetYamlConfig yaml configuration of Dashboard binary.
func (self *configBuilder) SetYamlConfig() *configBuilder {
	yamlFile, err := ioutil.ReadFile(Holder.GetConfigYamlPath())
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, self.config)
	if err != nil {
		log.Fatal(err)
	}
	return self
}

// GetConfigBuilder returns singleton instance of argument config builder.
func GetConfigBuilder() *configBuilder {
	return builderConf
}
