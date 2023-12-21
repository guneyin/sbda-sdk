package sdk

import (
	proto "github.com/guneyin/sbda-sdk/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (ds *DiscoveryService) GetServiceConn(svc ServiceEnum) (*grpc.ClientConn, error) {
	serviceInfo, err := ds.GetServiceInfo(svc.String())
	if err != nil {
		return nil, err
	}

	return grpc.Dial(serviceInfo.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func (ds *DiscoveryService) GetAuthServiceClient(conn *grpc.ClientConn) (proto.AuthServiceClient, error) {
	return proto.NewAuthServiceClient(conn), nil
}
