package scrapers

import (
	"fmt"
	"time"

	"github.com/derarken/vlr-api/src/utils"
	"github.com/gocolly/colly"
)

func GetLocation() (*time.Location, error) {
	c := colly.NewCollector()

	var abbreviation string
	c.OnHTML(".h-match-preview-time", func(e *colly.HTMLElement) {
		if abbreviation != "" {
			return
		}
		t := utils.PrettifyString(e.Text)
		abbreviation = t[len(t)-4:]
	})

	c.Visit("https://vlr.gg/")

	location, err := abbreviationToLocation(abbreviation)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func abbreviationToLocation(abbreviation string) (*time.Location, error) {
	zoneMap := map[string]string{
		"WEST": "WET",
		"PDT":  "PST",
		"NZDT": "NZST",
		"NDT":  "NST",
		"MEST": "MET",
		"MDT":  "MST",
		"IDT":  "IST",
		"HDT":  "HST",
		"EEST": "EET",
		"EDT":  "EST",
		"CEST": "CET",
		"CDT":  "CST",
		"BST":  "GMT",
		"AKDT": "AKST",
		"AEDT": "AEST",
		"ADT":  "AST",
		"ACDT": "ACST",
	}

	locationName, ok := zoneMap[abbreviation]
	if !ok {
		return nil, fmt.Errorf("unknown time zone abbreviation: %s", abbreviation)
	}

	location, err := time.LoadLocation(locationName)
	if err != nil {
		return nil, err
	}

	return location, nil
}
