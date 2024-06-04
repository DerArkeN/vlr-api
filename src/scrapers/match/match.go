package scraper_match

import (
	"strings"

	"github.com/derarken/vlr-api/src/customErrors"
	"github.com/derarken/vlr-api/src/utils"
	"github.com/gocolly/colly"
)

type Match struct {
	Super           *Super   `selector:".match-header > .match-header-super"`
	Versus          *Versus  `selector:".match-header > .match-header-vs"`
	StreamsTwitch   []string `selector:".match-streams-bets-container > .match-streams > .match-streams-container > .wf-card > a" attr:"href"`
	StreamsExternal []string `selector:".match-streams-bets-container > .match-streams > .match-streams-container > a" attr:"href"`
	Vods            []string `selector:".match-streams-bets-container > .match-vods > .match-streams-container > a" attr:"href"`
	// Maps            []*Map   `selector:".wf-card > .vm-stats > div > .vm-stats-gamesnav-container > *"`
	Maps []*Map `selector:".wf-card > .vm-stats > .vm-stats-container > .vm-stats-game"`
}

type Super struct {
	EventId string `selector:"div > .match-header-event" attr:"href"`
	// UTC Datetime string
	DateTime string `selector:"div > .match-header-date > .moment-tz-convert" attr:"data-utc-ts"`
	Patch    string `selector:"div > .match-header-date > [style='margin-top: 4px;']"`
}

type Versus struct {
	Team1Id string `selector:".match-header-link.mod-1" attr:"href"`
	Team2Id string `selector:".match-header-link.mod-2" attr:"href"`
	// the maps score in format "{Team1Score}:{Team2Score}"
	Score string `selector:".match-header-vs-score > .match-header-vs-score"`
	// live, final or the time until the match
	Notes []string `selector:".match-header-vs-note"`
}

type Map struct {
	Index  string
	Name   string   `selector:".vm-stats-game-header > .map > div > span"`
	Rounds []string `selector:"div > div > .vlr-rounds > .vlr-rounds-row > .vlr-rounds-row-col" attr:"title"`
}

func ScrapeMatchDetail(id string) (*Match, error) {
	c := colly.NewCollector()

	matchDetail := &Match{}
	c.OnHTML(".col.mod-3", func(e *colly.HTMLElement) {
		e.Unmarshal(matchDetail)
		matchDetail.Versus.Score = utils.PrettifyString(matchDetail.Versus.Score)
		matchDetail.Versus.Score = strings.ReplaceAll(matchDetail.Versus.Score, "vs.", "")

		for _, map_ := range matchDetail.Maps {
			map_.Name = utils.PrettifyString(map_.Name)
			map_.Name = strings.ReplaceAll(map_.Name, "PICK", "")
		}

		// remove all stats
		for i, map_ := range matchDetail.Maps {
			if map_.Name == "" {
				matchDetail.Maps = append(matchDetail.Maps[:i], matchDetail.Maps[i+1:]...)
			}
		}

		for i, map_ := range matchDetail.Maps {
			if map_.Name == "" {
				matchDetail.Maps = append(matchDetail.Maps[:i], matchDetail.Maps[i+1:]...)
			}
		}

		for i, map_ := range matchDetail.Maps {
			var filteredRounds []string
			for _, round := range map_.Rounds {
				if round != "" {
					filteredRounds = append(filteredRounds, round)
				}
			}
			matchDetail.Maps[i].Rounds = filteredRounds
		}
	})

	c.Visit("https://vlr.gg/" + id)

	if matchDetail == nil {
		return nil, customErrors.ErrNoMatch
	}

	return matchDetail, nil
}
