package api

import (
	"errors"
	"time"

	"github.com/derarken/vlr-api/proto"
	scraper_location "github.com/derarken/vlr-api/src/scraper/location"
	scraper_matches "github.com/derarken/vlr-api/src/scraper/matches"
)

type vlrStatus string

const (
	VLR_STATUS_LIVE      vlrStatus = "LIVE"
	VLR_STATUS_UPCOMING  vlrStatus = "Upcoming"
	VLR_STATUS_COMPLETED vlrStatus = "Completed"
)

var (
	ErrFromAfterTo = errors.New("from time must not be after to time")
	ErrToInFuture  = errors.New("to time must not be in the future when status is completed")
)

func GetMatchIds(status proto.Status, from time.Time, to time.Time) ([]string, error) {
	if from.After(to) {
		return nil, ErrFromAfterTo
	}

	switch status {
	case proto.Status_STATUS_LIVE:
		return getUpcomingMatchIds(from, to, VLR_STATUS_LIVE)
	case proto.Status_STATUS_UPCOMING:
		return getUpcomingMatchIds(from, to, VLR_STATUS_UPCOMING)
	case proto.Status_STATUS_COMPLETED:
		if to.After(time.Now()) {
			return nil, ErrToInFuture
		}
		return getCompletedMatchIds(from, to)
	}

	return nil, errors.New("invalid status")
}

func getUpcomingMatchIds(from time.Time, to time.Time, vlrStatus vlrStatus) ([]string, error) {
	var ids []string

	loc, err := scraper_location.GetLocation()
	if err != nil {
		return nil, err
	}

	var matches []*scraper_matches.Match
	page := 1
	for {
		newMatches, err := scraper_matches.ScrapeMatches(page)
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

	for _, match := range matches {
		if match.Status.Status != string(vlrStatus) {
			continue
		}

		matchTime, err := match.GetUtcTime(loc)
		if err != nil {
			return nil, err
		}

		if matchTime.UnixMilli() >= from.UnixMilli() && matchTime.UnixMilli() <= to.UnixMilli() {
			ids = append(ids, match.MatchId)
		}
	}

	return ids, nil
}

func getCompletedMatchIds(from time.Time, to time.Time) ([]string, error) {
	var ids []string

	loc, err := scraper_location.GetLocation()
	if err != nil {
		return nil, err
	}

	var matches []*scraper_matches.Match
	page := 1
	for {
		newMatches, err := scraper_matches.ScrapeResults(page)
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

	for _, match := range matches {
		if match.Status.Status != string(VLR_STATUS_COMPLETED) {
			continue
		}

		matchTime, err := match.GetUtcTime(loc)
		if err != nil {
			return nil, err
		}

		if matchTime.UnixMilli() >= from.UnixMilli() && matchTime.UnixMilli() <= to.UnixMilli() {
			ids = append(ids, match.MatchId)
		}
	}

	return ids, nil
}
