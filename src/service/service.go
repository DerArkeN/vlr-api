package service

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
	"github.com/flowchartsman/swaggerui"
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

func Start(swaggerJSON []byte) {
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

	startGateway(swaggerJSON)
}

func startGateway(swaggerJSON []byte) {
	conn, err := grpc.NewClient(fmt.Sprintf("dns:///%s", grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	gatewayMux := runtime.NewServeMux()
	err = proto.RegisterApiHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr: gatewayPort,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/v1") {
				gatewayMux.ServeHTTP(w, r)
				return
			}
			swaggerui.Handler(swaggerJSON).ServeHTTP(w, r)
		}),
	}

	grpclog.Infof("gRPC gateway %s and SwaggerUI running on %s", proto.Api_ServiceDesc.ServiceName, gatewayPort)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
