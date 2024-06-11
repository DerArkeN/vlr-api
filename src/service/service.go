package service

import (
	"context"
	"net"
	"net/http"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort    = ":8080"
	gatewayPort = ":8090"
)

type Server struct {
	proto.UnimplementedApiServer
}

func NewServer() *Server {
	return &Server{}
}

func Start() {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	apiServer := NewServer()
	proto.RegisterApiServer(s, apiServer)

	reflection.Register(s)

	grpclog.Infof("gRPC service %s running on %s", proto.Api_ServiceDesc.ServiceName, grpcPort)
	go func() {
		err = s.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	startGateway()
}

func startGateway() {
	conn, err := grpc.NewClient(grpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	mux := runtime.NewServeMux()
	err = proto.RegisterApiHandler(context.Background(), mux, conn)
	if err != nil {
		panic(err)
	}

	grpclog.Infof("gRPC gateway %s running on %s", proto.Api_ServiceDesc.ServiceName, gatewayPort)
	err = http.ListenAndServe(gatewayPort, mux)
	if err != nil {
		panic(err)
	}
}
