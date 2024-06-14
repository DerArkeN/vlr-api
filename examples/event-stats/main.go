package main

import (
	"context"
	"fmt"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := proto.NewApiClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := c.GetMatchIds(ctx, &proto.GetMatchIdsRequest{
		Status: proto.Status_STATUS_COMPLETED,
		Options: &proto.GetMatchIdsRequest_Options{
			EventId: "1999", // Champions Tour 2024: Masters Shanghai
		},
	})
	if err != nil {
		panic(err)
	}

	ids := resp.MatchIds

	matches := make([]*proto.Match, 0, len(ids))
	for i, id := range ids {
		// exclude showmatch
		if id == "359463" {
			continue
		}

		resp, err := c.GetMatch(ctx, &proto.GetMatchRequest{MatchId: id})
		if err != nil {
			panic(err)
		}

		matches = append(matches, resp.Match)
		fmt.Printf("Fetched match %d/%d\n", i+1, len(ids))
	}

	amount := 0
	for _, match := range matches {
		if teamWonSeriesAfterFirstMap(match) {
			amount++
		}
	}

	percentage := float64(amount) / float64(len(matches)) * 100
	fmt.Printf("Amount of matches where the team won the series after the first map: %d/%d (%.2f%%)\n", amount, len(matches), percentage)
}

func teamWonSeriesAfterFirstMap(match *proto.Match) bool {
	lastRoundFirstMap := match.Maps[0].Rounds[len(match.Maps[0].Rounds)-1]

	firstMapWonTeam1 := lastRoundFirstMap.Score1 > lastRoundFirstMap.Score2
	seriesWonTeam1 := match.Versus.Score1 > match.Versus.Score2

	if firstMapWonTeam1 && seriesWonTeam1 {
		return true
	}

	if !firstMapWonTeam1 && !seriesWonTeam1 {
		return true
	}

	return false
}
