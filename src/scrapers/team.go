package scrapers

import (
	"errors"
	"strings"

	"github.com/gocolly/colly"
)

var (
	ErrNoTeam = errors.New("no team found")
)

type Team struct {
	Name    string         `selector:".wf-card.mod-header.mod-full > .team-header > .team-header-desc > div > .team-header-name > h1"`
	Tricode string         `selector:".wf-card.mod-header.mod-full > .team-header > .team-header-desc > div > .team-header-name > h2"`
	Region  string         `selector:".wf-card.mod-header.mod-full > .team-header > .team-header-desc > div > .team-header-country"`
	Rating  string         `selector:".team-summary-container > .team-summary-container-1 > .wf-card.mod-rating > .team-core-block.mod-active > .core-rating-block.mod-active > .team-rating-info > .team-rating-info-section.mod-rating > .rating-num"`
	Players []*RosterEntry `selector:".team-summary-container > .team-summary-container-1 > :nth-child(9) > :nth-child(2) > .team-roster-item"`
	Staff   []*RosterEntry `selector:".team-summary-container > .team-summary-container-1 > :nth-child(9) > :nth-child(4) > .team-roster-item"`
}

type RosterEntry struct {
	PlayerID   string `selector:"a" attr:"href"`
	PlayerName string `selector:"a > .team-roster-item-name > .team-roster-item-name-alias"`
	RealName   string `selector:"a > .team-roster-item-name > .team-roster-item-name-real"`
	// "" for Player or Manager/Head Coach/Assistent Coach/etc.
	Role string `selector:"a > .team-roster-item-name > .wf-tag.mod-light"`
}

func ScrapeTeam(id string) (*Team, error) {
	if id == "" {
		return nil, ErrNoTeam
	}

	c := colly.NewCollector()

	var team *Team
	c.OnHTML(".col.mod-1", func(e *colly.HTMLElement) {
		team = &Team{}
		e.Unmarshal(team)

		team.trimPlayerIds()
	})

	c.Visit("https://www.vlr.gg/team/" + id)

	return team, nil
}

func (t *Team) trimPlayerIds() {
	for _, player := range t.Players {
		player.PlayerID = strings.Split(player.PlayerID, "/")[2]
	}

	for _, staff := range t.Staff {
		staff.PlayerID = strings.Split(staff.PlayerID, "/")[2]
	}
}
