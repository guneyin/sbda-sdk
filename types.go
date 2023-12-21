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

type ServiceProtocol int

const (
	ServiceProtocolHttp = iota
	ServiceProtocolGrpc
)

func (sn ServiceName) String() string {
	return string(sn)
}

type ServiceInfo struct {
	ID        string
	Name      string
	IP        string
	Port      int
	Protocol  ServiceProtocol
	CheckAddr string
}

func (si *ServiceInfo) Address() string {
	return fmt.Sprintf("%s:%d", si.IP, si.Port)
}
