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
	consul *consulapi.Client
}

func NewDiscoveryService(addr string) (*DiscoveryService, error) {
	consul, err := consulapi.NewClient(&consulapi.Config{Address: addr})
	if err != nil {
		return nil, err
	}

	return &DiscoveryService{consul: consul}, nil
}

func (ds *DiscoveryService) GetServiceInfo(serviceName string) (*ServiceInfo, error) {
	services, err := ds.consul.Agent().Services()
	if err != nil {
		return nil, err
	}

	ch, _ := ds.consul.Agent().Checks()
	for _, check := range ch {
		fmt.Println(check)
	}

	if service, ok := services[serviceName]; ok {
		addrs, _, err := ds.consul.Health().Service(serviceName, "", true, nil)
		if len(addrs) == 0 && err == nil {
			return nil, fmt.Errorf("service ( %s ) was not found", service.Service)
		}
		if err != nil {
			return nil, err
		}

		svc := addrs[0].Service

		return &ServiceInfo{
			ID:   svc.ID,
			Name: svc.Service,
			IP:   svc.Address,
			Port: svc.Port,
		}, nil
	}

	return nil, errors.New("service not found")
}

func (ds *DiscoveryService) RegisterService(svc IService) error {
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

	err := ds.consul.Agent().ServiceRegister(registration)
	if err != nil {
		return fmt.Errorf("failed to register service: %s - %s", si.Name, err.Error())
	}

	fmt.Printf("successfully register service: %s - %s", si.Name, si.Address())

	return nil
}

func (ds *DiscoveryService) DeRegisterService(svc IService) error {
	return ds.consul.Agent().ServiceDeregister(svc.ServiceInfo().ID)
}
