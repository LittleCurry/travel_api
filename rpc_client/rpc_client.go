package rpc_client

import (
	"errors"
	"log"
	"net"
	"time"

	pb "git.iguiyu.com/park/struct/protobuf"
	"git.iguiyu.com/park/api/configs"
	"github.com/alexfans/grpcpool"
	"google.golang.org/grpc"
)

var (
	smsGatewayPool *grpcpool.ConnectionPool
)

func Dialer(target string, timeout time.Duration) (net.Conn, error) {
	log.Println("connection .......")
	var (
		conn net.Conn
		err  error
	)
	for i := 1; i <= 4; i++ {
		conn, err = net.Dial("tcp", target)
		if err == nil {
			return conn, err
		}
		log.Printf("[%d] %s connection faild", i, target)
		time.Sleep(time.Duration(i) * time.Second)
	}
	panic(errors.New(target + "connect timeout"))
	return conn, err
}

func InitRpcConn() {
	var err error

	smsGatewayPool, err = grpcpool.NewConnectionPool(1, func() (*grpc.ClientConn, error) {
		log.Printf("SmsGatewayServer is [%s]", configs.AppConf.SmsGatewayServer)
		conn, err := grpc.Dial(configs.AppConf.SmsGatewayServer, grpc.WithInsecure(), grpc.FailOnNonTempDialError(false), grpc.WithDialer(Dialer))
		return conn, err
	})
	if err != nil {
		panic(err)
	}

}

func GetSmsGatewayClient() pb.SmsGatewayRpcClient {
	connection, err := smsGatewayPool.Get()
	defer connection.Close()
	if err != nil {
		log.Println("error is %s", err)
	}
	c := pb.NewSmsGatewayRpcClient(connection.Get())
	return c
}
