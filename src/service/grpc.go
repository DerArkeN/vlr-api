package service

import (
	"context"
	"net"

	"github.com/derarken/vlr-api/proto"
	"github.com/derarken/vlr-api/src/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Server struct {
	proto.UnimplementedApiServer
}

func StartGrpc() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterApiServer(s, &Server{})

	grpclog.Infof("Service %s running on %s", proto.Api_ServiceDesc.ServiceName, listener.Addr())
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *Server) GetMatchIds(ctx context.Context, in *proto.GetMatchIdsRequest) (*proto.GetMatchIdsResponse, error) {
	ids, err := api.GetMatchIds(in.Status, in.From.AsTime(), in.To.AsTime())
	if err != nil {
		return nil, err
	}
	return &proto.GetMatchIdsResponse{MatchIds: ids}, nil
}

func (s *Server) GetMatch(ctx context.Context, in *proto.GetMatchRequest) (*proto.GetMatchResponse, error) {
	match, err := api.GetMatch(in.MatchId)
	if err != nil {
		return nil, err
	}
	return &proto.GetMatchResponse{Match: match}, nil
}
