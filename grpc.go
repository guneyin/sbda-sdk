package sdk

import (
	auth "github.com/guneyin/sbda-proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func dialService(info *ServiceInfo) (*grpc.ClientConn, error) {
	return grpc.Dial(info.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func (ds *DiscoveryService) GetAuthService() (auth.AuthServiceClient, error) {
	serviceInfo, err := ds.GetServiceInfo(AuthServiceName.String())
	if err != nil {
		return nil, err
	}

	conn, err := dialService(serviceInfo)
	if err != nil {
		return nil, err
	}

	return auth.NewAuthServiceClient(conn), nil
}
