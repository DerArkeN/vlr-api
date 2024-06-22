package api

import (
	"errors"
	"fmt"
	"time"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
	"github.com/derarken/vlr-api/src/scrapers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrFromAfterTo = errors.New("from time must not be after to time")
	ErrToInFuture  = errors.New("to time must not be in the future when state is completed")
)

func GetMatchIds(request *proto.GetMatchIdsRequest) ([]string, error) {
	switch request.State {
	case proto.MatchState_MATCH_STATE_LIVE:
		return getLiveMatches(request.Options)

	case proto.MatchState_MATCH_STATE_UPCOMING:
		from, to, err := validateUpcomingTimes(request.From, request.To)
		if err != nil {
			return nil, err
		}
		return getUpcomingMatchIds(from, to, request.Options)

	case proto.MatchState_MATCH_STATE_COMPLETED:
		from, to, err := validateCompletedTimes(request.From, request.To)
		if err != nil {
			return nil, err
		}
		return getCompletedMatchIds(from, to, request.Options)

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

func getLiveMatches(opt *proto.GetMatchIdsRequest_Options) ([]string, error) {
	var matchIds []*scrapers.MatchIds
	page := 1
	for {
		var newMatchIds []*scrapers.MatchIds
		var err error
		if opt != nil && opt.EventId != "" {
			newMatchIds, err = scrapers.ScrapeEventMatches(opt.EventId)
			if err != nil {
				return nil, err
			}
			matchIds = append(matchIds, newMatchIds...)
			break
		} else {
			newMatchIds, err = scrapers.ScrapeMatches(page)
		}
		if err == scrapers.ErrNoMatchIds {
			break
		}
		if err != nil {
			return nil, err
		}
		matchIds = append(matchIds, newMatchIds...)

		lastMatchId := newMatchIds[len(newMatchIds)-1]
		if lastMatchId.State.State != string(scrapers.VLR_STATE_LIVE) {
			break
		}

		page++
	}

	return getIdsByStateAndTime(matchIds, scrapers.VLR_STATE_LIVE, nil, time.Time{}, time.Time{})
}

func validateUpcomingTimes(frompb *timestamppb.Timestamp, topb *timestamppb.Timestamp) (time.Time, time.Time, error) {
	from := time.Time{}
	if frompb != nil {
		from = frompb.AsTime()
	}
	to := time.Time{}
	if topb != nil {
		to = topb.AsTime()
	}

	if from.IsZero() && to.IsZero() {
		from = time.Now()
		to = from.Add(time.Hour * 24)
	}
	if from.IsZero() && !to.IsZero() {
		from = time.Now()
	}
	if !from.IsZero() && to.IsZero() {
		to = from.Add(time.Hour * 24)
	}
	if from.After(to) {
		return time.Time{}, time.Time{}, ErrFromAfterTo
	}
	return from, to, nil
}

func getUpcomingMatchIds(from time.Time, to time.Time, opt *proto.GetMatchIdsRequest_Options) ([]string, error) {
	loc, err := scrapers.GetLocation()
	if err != nil {
		return nil, err
	}

	var matchIds []*scrapers.MatchIds
	page := 1
	for {
		var newMatchIds []*scrapers.MatchIds
		var err error
		if opt != nil && opt.EventId != "" {
			newMatchIds, err = scrapers.ScrapeEventMatches(opt.EventId)
			if err != nil {
				return nil, err
			}
			matchIds = append(matchIds, newMatchIds...)
			from = time.Time{}
			to = time.Time{}
			break
		} else {
			newMatchIds, err = scrapers.ScrapeMatches(page)
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

	return getIdsByStateAndTime(matchIds, scrapers.VLR_STATE_UPCOMING, loc, from, to)
}

func validateCompletedTimes(frompb *timestamppb.Timestamp, topb *timestamppb.Timestamp) (time.Time, time.Time, error) {
	from := time.Time{}
	if frompb != nil {
		from = frompb.AsTime()
	}
	to := time.Time{}
	if topb != nil {
		to = topb.AsTime()
	}

	if from.IsZero() && to.IsZero() {
		to = time.Now()
		from = to.Add(time.Hour * -24)
	}
	if from.IsZero() && !to.IsZero() {
		from = to.Add(time.Hour * -24)
	}
	if !from.IsZero() && to.IsZero() {
		to = time.Now()
	}
	if from.After(to) {
		return time.Time{}, time.Time{}, ErrFromAfterTo
	}
	if to.After(time.Now()) {
		return time.Time{}, time.Time{}, ErrToInFuture
	}
	return from, to, nil
}

func getCompletedMatchIds(from time.Time, to time.Time, opt *proto.GetMatchIdsRequest_Options) ([]string, error) {
	loc, err := scrapers.GetLocation()
	if err != nil {
		return nil, err
	}

	var matchIds []*scrapers.MatchIds
	page := 1
	for {
		var newMatchIds []*scrapers.MatchIds
		var err error
		if opt != nil && opt.EventId != "" {
			newMatchIds, err = scrapers.ScrapeEventMatches(opt.EventId)
			if err != nil {
				return nil, err
			}
			matchIds = append(matchIds, newMatchIds...)
			from = time.Time{}
			to = time.Time{}
			break
		} else {
			newMatchIds, err = scrapers.ScrapeResults(page)
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

	return getIdsByStateAndTime(matchIds, scrapers.VLR_STATE_COMPLETED, loc, from, to)
}

func getIdsByStateAndTime(matches []*scrapers.MatchIds, state scrapers.VlrState, loc *time.Location, from time.Time, to time.Time) ([]string, error) {
	var ids []string

	for _, match := range matches {
		if match.State.State != string(state) {
			continue
		}

		if shouldValidateTime(loc, from, to) {
			matchTime, err := match.GetUtcTime(loc)
			if err != nil {
				return nil, err
			}

			if matchTime.UnixMilli() >= from.UnixMilli() && matchTime.UnixMilli() <= to.UnixMilli() {
				ids = append(ids, match.MatchId)
			}
		} else {
			ids = append(ids, match.MatchId)
		}

	}

	return ids, nil
}

func shouldValidateTime(loc *time.Location, from time.Time, to time.Time) bool {
	return loc != nil && !from.IsZero() && !to.IsZero()
}
