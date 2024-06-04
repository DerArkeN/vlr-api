package api

import (
	"errors"
	"time"

	"github.com/derarken/vlr-api/proto"
	"github.com/derarken/vlr-api/src/customErrors"
	scraper_location "github.com/derarken/vlr-api/src/scraper/location"
	scraper_matches "github.com/derarken/vlr-api/src/scraper/matches"
)

const (
	VLR_STATUS_LIVE     = "LIVE"
	VLR_STATUS_UPCOMING = "Upcoming"
)

func GetMatchIds(status proto.Status, from time.Time, to time.Time) ([]string, error) {
	switch status {
	case proto.Status_STATUS_LIVE:
		return getLiveMatchIds()
	case proto.Status_STATUS_UPCOMING:
		return getUpcomingMatchIds(from, to)
	}

	return nil, errors.New("invalid status")
}

func getLiveMatchIds() ([]string, error) {
	var ids []string

	matches, err := scraper_matches.ScrapeMatches(1)
	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		if match.Status.Status != VLR_STATUS_LIVE {
			continue
		}

		ids = append(ids, match.MatchId)
	}

	return ids, nil
}

func getUpcomingMatchIds(from time.Time, to time.Time) ([]string, error) {
	var ids []string

	loc, err := scraper_location.GetLocation()
	if err != nil {
		return nil, err
	}

	var matches []*scraper_matches.Match
	page := 1
	for {
		newMatches, err := scraper_matches.ScrapeMatches(page)
		if err == customErrors.ErrNoMatches {
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
		if match.Status.Status != VLR_STATUS_UPCOMING {
			continue
		}

		matchTime, err := match.GetUtcTime(loc)
		if err != nil {
			return nil, err
		}

		if matchTime.After(from) && matchTime.Before(to) {
			ids = append(ids, match.MatchId)
		}
	}

	return ids, nil
}
