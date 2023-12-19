package sdk

import (
	"fmt"
)

type ServiceName string

const (
	ConsulServiceName ServiceName = "consul-service"
	AuthServiceName   ServiceName = "auth-service"
	ApilServiceName   ServiceName = "api-service"
)

func (sn ServiceName) String() string {
	return string(sn)
}

type ServiceInfo struct {
	Name string
	IP   string
	Port int
}

func (si *ServiceInfo) Address() string {
	return fmt.Sprintf("%s:%d", si.IP, si.Port)
}
