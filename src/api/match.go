package api

import (
	"strconv"
	"strings"

	"github.com/derarken/vlr-api/proto"
	scraper_match "github.com/derarken/vlr-api/src/scraper/match"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetMatch(matchId string) (*proto.Match, error) {
	smatch, err := scraper_match.ScrapeMatchDetail(matchId)
	if err != nil {
		return nil, err
	}

	utcTime, err := smatch.GetUtcTime()
	if err != nil {
		return nil, err
	}

	score1, score2, err := smatch.GetScore()
	if err != nil {
		return nil, err
	}

	maps, err := getMaps(smatch.Maps)
	if err != nil {
		return nil, err
	}

	match := &proto.Match{
		Head: &proto.Match_Head{
			MatchId:  matchId,
			EventId:  smatch.Super.EventId,
			DateTime: timestamppb.New(utcTime),
		},
		Versus: &proto.Match_Versus{
			Team1: &proto.Team{
				TeamId: smatch.Versus.Team1Id,
				Name:   smatch.Versus.Team1Name,
			},
			Score1: int32(score1),
			Team2: &proto.Team{
				TeamId: smatch.Versus.Team2Id,
				Name:   smatch.Versus.Team2Name,
			},
			Score2: int32(score2),
		},
		Maps: maps,
	}

	return match, nil
}

func getMaps(maps []*scraper_match.Map) ([]*proto.Match_Map, error) {
	var protoMaps []*proto.Match_Map
	for _, map_ := range maps {
		protoRounds, err := getRounds(map_.Rounds)
		if err != nil {
			return nil, err
		}
		protoMaps = append(protoMaps, &proto.Match_Map{
			Name:   map_.Name,
			Rounds: protoRounds,
		})
	}
	return protoMaps, nil
}

func getRounds(rounds []string) ([]*proto.Match_Map_Round, error) {
	var protoRounds []*proto.Match_Map_Round
	for _, round := range rounds {
		scores := strings.Split(round, "-")
		score1, err := strconv.Atoi(scores[0])
		if err != nil {
			return nil, err
		}
		score2, err := strconv.Atoi(scores[1])
		if err != nil {
			return nil, err
		}
		protoRounds = append(protoRounds, &proto.Match_Map_Round{
			Score1: int32(score1),
			Score2: int32(score2),
		})
	}
	return protoRounds, nil
}
