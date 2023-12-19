package sdk

import (
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
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
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	si := svc.ServiceInfo()

	registration := &consulapi.AgentServiceRegistration{
		ID:      si.Name,
		Name:    si.Name,
		Port:    si.Port,
		Address: si.IP,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s/check", si.Address()),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regiErr := consul.Agent().ServiceRegister(registration)

	if regiErr != nil {
		return fmt.Errorf("failed to register service: %s", si.Name)
	}

	fmt.Printf("successfully register service: %s", si.Name)

	return nil
}
