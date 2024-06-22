package scrapers

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/derarken/vlr-api/src/utils"
	"github.com/gocolly/colly"
)

var (
	ErrNoMatch      = errors.New("no match found")
	ErrInvalidScore = errors.New("invalid score")
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
	EventId   string `selector:"div > .match-header-event" attr:"href"`
	EventName string `selector:"div > .match-header-event > div > [style='font-weight: 700;']"`
	Stage     string `selector:"div > .match-header-event > div > .match-header-event-series"`
	// this is not UTC but America/Havana
	DateTime string `selector:"div > .match-header-date > .moment-tz-convert" attr:"data-utc-ts"`
	Patch    string `selector:"div > .match-header-date > [style='margin-top: 4px;']"`
}

type Versus struct {
	Team1Name string `selector:".match-header-link.mod-1 > .match-header-link-name.mod-1 > .wf-title-med"`
	Team1Id   string `selector:".match-header-link.mod-1" attr:"href"`
	Team2Name string `selector:".match-header-link.mod-2 > .match-header-link-name.mod-2 > .wf-title-med"`
	Team2Id   string `selector:".match-header-link.mod-2" attr:"href"`
	// the maps score in format "{Team1Score}:{Team2Score}"
	Score string `selector:".match-header-vs-score > .match-header-vs-score"`
	// live, final or the time until the match
	Notes []string `selector:".match-header-vs-note"`
}

type Map struct {
	Name   string   `selector:".vm-stats-game-header > .map > div > span"`
	Rounds []string `selector:"div > div > .vlr-rounds > .vlr-rounds-row > .vlr-rounds-row-col" attr:"title"`
}

func ScrapeMatchDetail(id string) (*Match, error) {
	if id == "" {
		return nil, ErrNoMatch
	}

	c := colly.NewCollector()

	var matchDetail *Match
	c.OnHTML(".col.mod-3", func(e *colly.HTMLElement) {
		matchDetail = &Match{}
		e.Unmarshal(matchDetail)
		if matchDetail.Super == nil || matchDetail.Versus == nil {
			matchDetail = nil
			return
		}

		matchDetail.Super.EventId = strings.Split(matchDetail.Super.EventId, "/")[2]
		matchDetail.Super.EventName = utils.PrettifyString(matchDetail.Super.EventName)
		matchDetail.Super.Stage = utils.PrettifyString(matchDetail.Super.Stage)

		matchDetail.Versus.Score = utils.PrettifyString(matchDetail.Versus.Score)
		matchDetail.Versus.Score = strings.ReplaceAll(matchDetail.Versus.Score, "vs.", "")

		matchDetail.Versus.Team1Id = strings.Split(matchDetail.Versus.Team1Id, "/")[2]
		matchDetail.Versus.Team2Id = strings.Split(matchDetail.Versus.Team2Id, "/")[2]

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
		return nil, ErrNoMatch
	}

	return matchDetail, nil
}

func (m *Match) GetUtcTime() (time.Time, error) {
	loc, err := time.LoadLocation("America/Havana")
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.ParseInLocation(time.DateTime, m.Super.DateTime, loc)
	if err != nil {
		return time.Time{}, err
	}
	t = t.UTC()
	return t, nil
}

func (m *Match) GetScore() (int, int, error) {
	scores := strings.Split(m.Versus.Score, ":")
	if len(scores) != 2 {
		return 0, 0, ErrInvalidScore
	}
	score1, err := strconv.Atoi(scores[0])
	if err != nil {
		return 0, 0, err
	}
	score2, err := strconv.Atoi(scores[1])
	if err != nil {
		return 0, 0, err
	}
	return score1, score2, nil
}
