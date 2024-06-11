package service

import (
	"context"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
	"github.com/derarken/vlr-api/src/api"
)

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
