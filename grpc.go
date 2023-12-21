package sdk

import (
	proto "github.com/guneyin/sbda-sdk/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func dialService(info *ServiceInfo) (*grpc.ClientConn, error) {
	return grpc.Dial(info.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func (ds *DiscoveryService) GetAuthService() (proto.AuthServiceClient, error) {
	serviceInfo, err := ds.GetServiceInfo(AuthServiceName.String())
	if err != nil {
		return nil, err
	}

	conn, err := dialService(serviceInfo)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return proto.NewAuthServiceClient(conn), nil
}
