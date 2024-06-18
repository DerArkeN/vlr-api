package scraper_matches

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type VlrState string

const (
	VLR_STATE_LIVE      VlrState = "LIVE"
	VLR_STATE_UPCOMING  VlrState = "Upcoming"
	VLR_STATE_COMPLETED VlrState = "Completed"
)

var ErrNoMatches = errors.New("no matches found")

type Match struct {
	MatchId string
	State   *State `selector:".match-item-eta"`
	Date    string `selector:".match-item-date"`
	Time    string `selector:".match-item-time"`
	// Teams  []*Team `selector:".match-item-vs > .match-item-vs-team"`
	// Note   string  `selector:".match-item-note"`
	// Vods   []*Vod  `selector:".match-item-vod"`
	// Event  string  `selector:".match-item-event"`
	// Series string  `selector:".match-item-event > .match-item-event-series"`
}

// type Team struct {
// 	Name  string `selector:".match-item-vs-team-name"`
// 	Score string `selector:".match-item-vs-team-score"`
// }

type State struct {
	State string `selector:".ml > .ml-status"`
	ETA   string `selector:".ml > .ml-eta"`
}

// type Vod struct {
// 	Label string   `selector:".wf-module-label"`
// 	Tags  []string `selector:".wf-tag"`
// }

// Fetches https://vlr.gg/matches/results
func ScrapeResults(page int) ([]*Match, error) {
	return scrapeMatches("https://vlr.gg/matches/results/?page=" + fmt.Sprint(page))
}

// Fetches https://vlr.gg/matches
func ScrapeMatches(page int) ([]*Match, error) {
	return scrapeMatches("https://vlr.gg/matches/?page=" + fmt.Sprint(page))
}

func ScrapeEventMatches(eventId string) ([]*Match, error) {
	return scrapeMatches("https://www.vlr.gg/event/matches/" + eventId + "/?series_id=all&group=all")
}

func scrapeMatches(url string) ([]*Match, error) {
	c := colly.NewCollector()

	matches := []*Match{}

	c.OnHTML(".col", func(column *colly.HTMLElement) {
		var currDate string
		column.ForEach("div", func(_ int, colEntry *colly.HTMLElement) {
			if colEntry.Attr("class") == "wf-label mod-large" {
				dateStrings := strings.Split(colEntry.Text, "\t")
				for _, dateString := range dateStrings {
					if strings.Contains(dateString, ",") {
						currDate = dateString
						currDate = strings.ReplaceAll(currDate, "\n", "")
						currDate = strings.TrimSpace(currDate)
						break
					}
				}
			}

			if colEntry.Attr("class") == "wf-card" {
				colEntry.ForEach("a.match-item", func(_ int, matchEntry *colly.HTMLElement) {
					match := &Match{}
					err := matchEntry.Unmarshal(match)
					if err != nil {
						log.Println("Error unmarshalling match: ", err)
						return
					}

					// the date is outside of the match element
					match.Date = currDate

					// head attribute can't be unmarshalled
					match.MatchId = strings.Split(matchEntry.Attr("href"), "/")[1]
					if match.MatchId == "" {
						return
					}

					// // event needs to be split since it contains the stage as well
					// strings := strings.Split(match.Event, "\t")
					// match.Event = strings[len(strings)-1]

					matches = append(matches, match)
				})
			}
		})
	})

	err := c.Visit(url)

	if err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		return nil, ErrNoMatches
	}

	return matches, nil
}

func (m *Match) GetUtcTime(location *time.Location) (time.Time, error) {
	if m.Time == "TBD" {
		m.Time = "00:00 AM"
	}

	matchTime, err := time.ParseInLocation("Mon, January 2, 2006, 15:04 PM", m.Date+", "+m.Time, location)
	if err != nil {
		return time.Time{}, err
	}

	matchTime = matchTime.UTC()

	return matchTime, nil
}
