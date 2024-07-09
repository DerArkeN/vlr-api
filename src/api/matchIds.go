package api

import (
	"fmt"
	"time"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
	"github.com/derarken/vlr-api/src/scrapers"
	"github.com/derarken/vlr-api/src/utils"
)

type MatchIdFactory struct{}

var matchIdFactory = &MatchIdFactory{}

func GetMatchIds(request *proto.GetMatchIdsRequest) ([]string, error) {
	switch request.State {
	case proto.MatchState_MATCH_STATE_LIVE:
		return matchIdFactory.getLiveMatches(request.Options)

	case proto.MatchState_MATCH_STATE_UPCOMING:
		from, to, err := utils.ValidateUpcomingTimes(request.From, request.To)
		if err != nil {
			return nil, err
		}
		return matchIdFactory.getUpcomingMatchIds(from, to, request.Options)

	case proto.MatchState_MATCH_STATE_COMPLETED:
		from, to, err := utils.ValidateCompletedTimes(request.From, request.To)
		if err != nil {
			return nil, err
		}
		return matchIdFactory.getCompletedMatchIds(from, to, request.Options)

	default:
		validStates := []string{}
		for k, v := range proto.MatchState_name {
			if k == 0 {
				continue
			}
			validStates = append(validStates, v)
		}
		return nil, fmt.Errorf("invalid state, valid states are %v", validStates)

	}
}

func (f *MatchIdFactory) getLiveMatches(opt *proto.GetMatchIdsRequest_Options) ([]string, error) {
	var matchIds []*scrapers.MatchId
	page := 1
	for {
		var newMatchIds []*scrapers.MatchId
		var err error
		if opt != nil && opt.EventId != "" {
			newMatchIds, err = scrapers.ScrapeEventMatchIds(opt.EventId)
			if err != nil {
				return nil, err
			}
			matchIds = append(matchIds, newMatchIds...)
			break
		} else {
			newMatchIds, err = scrapers.ScrapeMatchIds(page)
		}
		if err == scrapers.ErrNoMatchIds {
			break
		}
		if err != nil {
			return nil, err
		}
		matchIds = append(matchIds, newMatchIds...)

		lastMatchId := newMatchIds[len(newMatchIds)-1]
		if lastMatchId.State.State != string(scrapers.MATCH_STATE_LIVE) {
			break
		}

		page++
	}

	return f.getMatchIdsByStateAndTime(matchIds, scrapers.MATCH_STATE_LIVE, nil, time.Time{}, time.Time{})
}

func (f *MatchIdFactory) getUpcomingMatchIds(from time.Time, to time.Time, opt *proto.GetMatchIdsRequest_Options) ([]string, error) {
	loc, err := scrapers.GetLocation()
	if err != nil {
		return nil, err
	}

	var matchIds []*scrapers.MatchId
	page := 1
	for {
		var newMatchIds []*scrapers.MatchId
		var err error
		if opt != nil && opt.EventId != "" {
			newMatchIds, err = scrapers.ScrapeEventMatchIds(opt.EventId)
			if err != nil {
				return nil, err
			}
			matchIds = append(matchIds, newMatchIds...)
			from = time.Time{}
			to = time.Time{}
			break
		} else {
			newMatchIds, err = scrapers.ScrapeMatchIds(page)
		}
		if err == scrapers.ErrNoMatchIds {
			break
		}
		if err != nil {
			return nil, err
		}
		matchIds = append(matchIds, newMatchIds...)

		lastMatchId := newMatchIds[len(newMatchIds)-1]
		matchTime, err := lastMatchId.GetUtcTime(loc)
		if err != nil {
			return nil, err
		}

		if matchTime.After(to) {
			break
		}

		page++
	}

	return f.getMatchIdsByStateAndTime(matchIds, scrapers.MATCH_STATE_UPCOMING, loc, from, to)
}

func (f *MatchIdFactory) getCompletedMatchIds(from time.Time, to time.Time, opt *proto.GetMatchIdsRequest_Options) ([]string, error) {
	loc, err := scrapers.GetLocation()
	if err != nil {
		return nil, err
	}

	var matchIds []*scrapers.MatchId
	page := 1
	for {
		var newMatchIds []*scrapers.MatchId
		var err error
		if opt != nil && opt.EventId != "" {
			newMatchIds, err = scrapers.ScrapeEventMatchIds(opt.EventId)
			if err != nil {
				return nil, err
			}
			matchIds = append(matchIds, newMatchIds...)
			from = time.Time{}
			to = time.Time{}
			break
		} else {
			newMatchIds, err = scrapers.ScrapeResultIds(page)
		}
		if err == scrapers.ErrNoMatchIds {
			break
		}
		if err != nil {
			return nil, err
		}
		matchIds = append(matchIds, newMatchIds...)

		lastMatchId := newMatchIds[len(newMatchIds)-1]
		matchTime, err := lastMatchId.GetUtcTime(loc)
		if err != nil {
			return nil, err
		}

		if matchTime.Before(to) {
			break
		}

		page++
	}

	return f.getMatchIdsByStateAndTime(matchIds, scrapers.MATCH_STATE_COMPLETED, loc, from, to)
}

func (f *MatchIdFactory) getMatchIdsByStateAndTime(matches []*scrapers.MatchId, state scrapers.MatchState, loc *time.Location, from time.Time, to time.Time) ([]string, error) {
	var ids []string

	for _, match := range matches {
		if match.State.State != string(state) {
			continue
		}

		if utils.ShouldValidateTime(loc, from, to) {
			matchTime, err := match.GetUtcTime(loc)
			if err != nil {
				return nil, err
			}

			if utils.IsTimeInRange(from, to, matchTime) {
				ids = append(ids, match.MatchId)
			}
		} else {
			ids = append(ids, match.MatchId)
		}

	}

	return ids, nil
}
