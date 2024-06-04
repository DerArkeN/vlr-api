package api

import (
	"time"

	scraper_matches "github.com/derarken/vlr-api/src/scrapers/matches"
)

const (
	MatchIdsPath = "/matchIds?status=%s&from=%s&to=%s"
)

type Status string

const (
	STATUS_LIVE      Status = "LIVE"
	STATUS_UPCOMING  Status = "UPCOMING"
	STATUS_COMPLETED Status = "COMPLETED"
)

type MatchIds struct {
	Ids []string `json:"ids"`
}

func GetMatchIds(status Status, from time.Time, to time.Time) (*MatchIds, error) {
	result := &MatchIds{}
	switch status {
	case STATUS_LIVE:
		matches, err := scraper_matches.ScrapeMatches(1)
		if err != nil {
			return nil, err
		}

		for _, match := range matches {
			if match.Status.Status != "LIVE" {
				continue
			}

			matchTime, err := match.ConvertMatchTime()
			if err != nil {
				return nil, err
			}

			if !((matchTime.After(from) || matchTime == from) && (matchTime.Before(to) || matchTime == to)) {
				continue
			}

			result.Ids = append(result.Ids, match.Id)
		}
	}

	return result, nil
}
