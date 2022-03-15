package args

import "net"

var Holder = &holder{}

// Argument holder structure. It is private to make sure that only 1 instance can be created. It holds all
// arguments values passed to Dashboard binary.
type holder struct {
	configYamlPath      string
	insecurePort        int
	insecureBindAddress net.IP
}

// GetConfigYamlPath 'config-yaml-path' argument of Dashboard binary.
func (self *holder) GetConfigYamlPath() string {
	return self.configYamlPath
}

// GetInsecurePort 'insecure-port' argument of Dashboard binary.
func (self *holder) GetInsecurePort() int {
	return self.insecurePort
}

// GetInsecureBindAddress 'insecure-bind-address' argument of Dashboard binary.
func (self *holder) GetInsecureBindAddress() net.IP {
	return self.insecureBindAddress
}
