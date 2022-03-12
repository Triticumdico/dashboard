package args

import "net"

var Holder = &holder{}

// Argument holder structure. It is private to make sure that only 1 instance can be created. It holds all
// arguments values passed to Dashboard binary.
type holder struct {
	insecurePort        int
	insecureBindAddress net.IP
}

// GetInsecurePort 'insecure-port' argument of Dashboard binary.
func (self *holder) GetInsecurePort() int {
	return self.insecurePort
}

// GetInsecureBindAddress 'insecure-bind-address' argument of Dashboard binary.
func (self *holder) GetInsecureBindAddress() net.IP {
	return self.insecureBindAddress
}
