package pkg

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateGRPCDial(serverAddr string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithTimeout(5 * time.Second))
	conn, err := grpc.Dial(serverAddr, opts...)
	return conn, err
}