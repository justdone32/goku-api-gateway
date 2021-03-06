package consul

import (
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
)

const DriverName = "consul"

func init() {

	discovery.RegisteredDiscovery(DriverName, discovery.NewDriver(Create))

}

func Create(config string) discovery.Discovery {
	return NewConsulDiscovery(config)
}
