package api

import (
	"strconv"
	"strings"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
	"github.com/derarken/vlr-api/src/scrapers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetMatch(matchId string) (*proto.Match, error) {
	smatch, err := scrapers.ScrapeMatchDetail(matchId)
	if err != nil {
		return nil, err
	}

	utcTime, err := smatch.GetUtcTime()
	if err != nil {
		return nil, err
	}

	score1 := 0
	score2 := 0
	score1, score2, err = smatch.GetScore()
	if err != nil && err != scrapers.ErrInvalidScore {
		return nil, err
	}

	maps, err := getMaps(smatch.Maps)
	if err != nil {
		return nil, err
	}

	match := &proto.Match{
		Head: &proto.Match_Head{
			State:   getProtoMatchState(smatch.Versus.Notes),
			MatchId: matchId,
			Event: &proto.Match_Head_Event{
				EventId: smatch.Super.EventId,
				Name:    smatch.Super.EventName,
				Stage:   smatch.Super.Stage,
			},
			DateTime: timestamppb.New(utcTime),
		},
		Versus: &proto.Match_Versus{
			Team1: &proto.Match_Versus_Team{
				TeamId: smatch.Versus.Team1Id,
				Name:   smatch.Versus.Team1Name,
			},
			Score1: int32(score1),
			Team2: &proto.Match_Versus_Team{
				TeamId: smatch.Versus.Team2Id,
				Name:   smatch.Versus.Team2Name,
			},
			Score2: int32(score2),
		},
		Maps: maps,
	}

	return match, nil
}

func getMaps(maps []*scrapers.Map) ([]*proto.Match_Map, error) {
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

func getProtoMatchState(notes []string) proto.MatchState {
	for _, note := range notes {
		if note == "live" {
			return proto.MatchState_MATCH_STATE_LIVE
		}
		if strings.Contains(note, "h") || strings.Contains(note, "m") {
			return proto.MatchState_MATCH_STATE_UPCOMING
		}
		if note == "final" || strings.Contains(note, "forfeited") {
			return proto.MatchState_MATCH_STATE_COMPLETED
		}
	}
	return proto.MatchState_MATCH_STATE_UNSPECIFIED
}
