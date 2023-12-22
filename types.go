package sdk

import (
	"fmt"
)

type ServiceEnum string

const (
	AuthService            ServiceEnum = "auth-service"
	ProductCategoryService ServiceEnum = "product-category-service"
)

type ServiceProtocol int

const (
	ServiceProtocolHttp = iota
	ServiceProtocolGrpc
)

func (s ServiceEnum) String() string {
	return string(s)
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
