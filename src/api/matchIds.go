package api

import (
	"errors"
	"fmt"
	"time"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
	scraper_location "github.com/derarken/vlr-api/src/scraper/location"
	scraper_matches "github.com/derarken/vlr-api/src/scraper/matches"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrFromAfterTo = errors.New("from time must not be after to time")
	ErrToInFuture  = errors.New("to time must not be in the future when status is completed")
)

func GetMatchIds(request *proto.GetMatchIdsRequest) ([]string, error) {
	switch request.Status {
	case proto.Status_STATUS_LIVE:
		return getLiveMatches(request.Options)

	case proto.Status_STATUS_UPCOMING:
		from, to, err := validateUpcomingTimes(request.From, request.To)
		if err != nil {
			return nil, err
		}
		return getUpcomingMatchIds(from, to, request.Options)

	case proto.Status_STATUS_COMPLETED:
		from, to, err := validateCompletedTimes(request.From, request.To)
		if err != nil {
			return nil, err
		}
		return getCompletedMatchIds(from, to, request.Options)

	default:
		validStatuses := []string{}
		for k, v := range proto.Status_name {
			if k == 0 {
				continue
			}
			validStatuses = append(validStatuses, v)
		}
		return nil, fmt.Errorf("invalid status, valid statuses are %v", validStatuses)

	}
}

func getLiveMatches(opt *proto.GetMatchIdsRequest_Options) ([]string, error) {
	var matches []*scraper_matches.Match
	page := 1
	for {
		var newMatches []*scraper_matches.Match
		var err error
		if opt != nil && opt.EventId != "" {
			newMatches, err = scraper_matches.ScrapeEventMatches(opt.EventId)
			if err != nil {
				return nil, err
			}
			matches = append(matches, newMatches...)
			break
		} else {
			newMatches, err = scraper_matches.ScrapeMatches(page)
		}
		if err == scraper_matches.ErrNoMatches {
			break
		}
		if err != nil {
			return nil, err
		}
		matches = append(matches, newMatches...)

		lastMatch := newMatches[len(newMatches)-1]
		if lastMatch.Status.Status != string(scraper_matches.VLR_STATUS_LIVE) {
			break
		}

		page++
	}

	return getIdsByStatusAndTime(matches, scraper_matches.VLR_STATUS_LIVE, nil, time.Time{}, time.Time{})
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
	loc, err := scraper_location.GetLocation()
	if err != nil {
		return nil, err
	}

	var matches []*scraper_matches.Match
	page := 1
	for {
		var newMatches []*scraper_matches.Match
		var err error
		if opt != nil && opt.EventId != "" {
			newMatches, err = scraper_matches.ScrapeEventMatches(opt.EventId)
			if err != nil {
				return nil, err
			}
			matches = append(matches, newMatches...)
			from = time.Time{}
			to = time.Time{}
			break
		} else {
			newMatches, err = scraper_matches.ScrapeMatches(page)
		}
		if err == scraper_matches.ErrNoMatches {
			break
		}
		if err != nil {
			return nil, err
		}
		matches = append(matches, newMatches...)

		lastMatch := newMatches[len(newMatches)-1]
		matchTime, err := lastMatch.GetUtcTime(loc)
		if err != nil {
			return nil, err
		}

		if matchTime.After(to) {
			break
		}

		page++
	}

	return getIdsByStatusAndTime(matches, scraper_matches.VLR_STATUS_UPCOMING, loc, from, to)
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
	loc, err := scraper_location.GetLocation()
	if err != nil {
		return nil, err
	}

	var matches []*scraper_matches.Match
	page := 1
	for {
		var newMatches []*scraper_matches.Match
		var err error
		if opt != nil && opt.EventId != "" {
			newMatches, err = scraper_matches.ScrapeEventMatches(opt.EventId)
			if err != nil {
				return nil, err
			}
			matches = append(matches, newMatches...)
			from = time.Time{}
			to = time.Time{}
			break
		} else {
			newMatches, err = scraper_matches.ScrapeResults(page)
		}
		if err == scraper_matches.ErrNoMatches {
			break
		}
		if err != nil {
			return nil, err
		}
		matches = append(matches, newMatches...)

		lastMatch := newMatches[len(newMatches)-1]
		matchTime, err := lastMatch.GetUtcTime(loc)
		if err != nil {
			return nil, err
		}

		if matchTime.Before(to) {
			break
		}

		page++
	}

	return getIdsByStatusAndTime(matches, scraper_matches.VLR_STATUS_COMPLETED, loc, from, to)
}

func getIdsByStatusAndTime(matches []*scraper_matches.Match, status scraper_matches.VlrStatus, loc *time.Location, from time.Time, to time.Time) ([]string, error) {
	var ids []string

	for _, match := range matches {
		if match.Status.Status != string(status) {
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
