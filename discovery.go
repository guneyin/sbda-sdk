package sdk

import (
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
)

type IService interface {
	Register() error
	UnRegister() error
	ServiceInfo() *ServiceInfo
}

type DiscoveryService struct {
	addr string
}

func NewDiscoveryService(addr string) *DiscoveryService {
	return &DiscoveryService{addr: addr}
}

func (ds *DiscoveryService) GetServiceInfo(serviceName string) (*ServiceInfo, error) {
	consul, err := consulapi.NewClient(&consulapi.Config{Address: ds.addr})
	if err != nil {
		return nil, err
	}

	services, err := consul.Agent().Services()
	if err != nil {
		return nil, err
	}

	if service, ok := services[serviceName]; ok {
		return &ServiceInfo{
			Name: serviceName,
			IP:   service.Address,
			Port: service.Port,
		}, nil
	}

	return nil, errors.New("service not found")
}

func (ds *DiscoveryService) RegisterService(svc IService) error {
	config := &consulapi.Config{Address: ds.addr}

	consul, err := consulapi.NewClient(config)
	if err != nil {
		return err
	}

	si := svc.ServiceInfo()

	registration := &consulapi.AgentServiceRegistration{
		ID:      si.Name,
		Name:    si.Name,
		Port:    si.Port,
		Address: si.IP,
		Check: &consulapi.AgentServiceCheck{
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	switch si.Protocol {
	case ServiceProtocolHttp:
		registration.Check.HTTP = fmt.Sprintf("%s/%s", si.Address(), si.CheckAddr)
	case ServiceProtocolGrpc:
		registration.Check.GRPC = fmt.Sprintf("%s", si.Address())
	default:
	}

	err = consul.Agent().ServiceRegister(registration)

	if err != nil {
		return fmt.Errorf("failed to register service: %s - %s", si.Name, err.Error())
	}

	fmt.Printf("successfully register service: %s - %s", si.Name, config.Address)

	return nil
}
