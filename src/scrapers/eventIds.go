package scrapers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/derarken/vlr-api/src/utils"
	"github.com/gocolly/colly"
)

const (
	upcomingContainer  = "#wrapper > div.col-container > div > div.events-container > div:nth-child(1)"
	completedContainer = "#wrapper > div.col-container > div > div.events-container > div:nth-child(2)"
)

var (
	ErrNoEventIds = fmt.Errorf("no events found")
)

type EventId struct {
	EventId string
	Title   string `selector:"div.event-item-inner > .event-item-title"`
	// ongoing, upcoming, completed
	State string `selector:"div.event-item-inner > div:nth-child(2) > div:nth-child(1) > span"`
	// only includes month and day (either "Jan 1-31", "Jan 1-Feb 2" or "Jan 1") since year is not provided (local time)
	DateRange string `selector:"div.event-item-inner > div:nth-child(2) > div.event-item-desc-item.mod-dates"`

	// these fields have to be filled manually if no year is provided in the year is set to 1970
	// can only be filled if at least one event on the list contains the year in the title
	Start time.Time
	End   time.Time
}

func ScrapeUpcomingOrOngoingEventIds(page int) ([]*EventId, error) {
	return scrapeEventIds(upcomingContainer, page)
}

func ScrapeCompletedEventIds(page int) ([]*EventId, error) {
	return scrapeEventIds(completedContainer, page)
}

func scrapeEventIds(containerQuery string, page int) ([]*EventId, error) {
	c := colly.NewCollector()

	eventIds := []*EventId{}
	c.OnHTML(containerQuery, func(container *colly.HTMLElement) {
		container.ForEach("a", func(_ int, eventEntry *colly.HTMLElement) {
			e := &EventId{}
			err := eventEntry.Unmarshal(e)
			if err != nil {
				return
			}

			e.EventId = strings.Split(eventEntry.Attr("href"), "/")[2]
			if e.EventId == "" {
				return
			}

			eventIds = append(eventIds, e)
		})
	})

	err := c.Visit("https://www.vlr.gg/events/?page=" + fmt.Sprint(page))
	if err != nil {
		return nil, err
	}

	if len(eventIds) == 0 {
		return nil, ErrNoEventIds
	}

	fillEventDates(eventIds)

	return eventIds, nil
}

func (e *EventId) getDayAndMonth() (start time.Time, end time.Time) {
	formatString := "Jan 2"

	dateRange := utils.PrettifyString(e.DateRange)
	dateRange = strings.TrimSuffix(dateRange, "Dates")

	dates := strings.Split(dateRange, "—")

	startDateString := dates[0]
	endDateString := startDateString
	if len(dates) == 2 && dates[1] != "TBD" {
		endDateString = dates[1]
	}

	startDate, err := time.Parse(formatString, startDateString)
	if err != nil {
		return time.Time{}, time.Time{}
	}

	endDate, err := time.Parse(formatString, endDateString)
	if err != nil {
		// this is neccecary because the format for the same month is "Jan 1-31" and for different months "Jan 1-Feb 2"
		newString := fmt.Sprintf("%s %s", startDateString[:3], endDateString)
		endDate, err = time.Parse(formatString, newString)
		if err != nil {
			return time.Time{}, time.Time{}
		}
	}

	return startDate, endDate
}

func (e *EventId) getYearFromTitle() int {
	re := regexp.MustCompile(`\d{4}`)
	yearString := re.FindString(e.Title)
	year, err := strconv.Atoi(yearString)
	const valorantRelease = 2020
	if err != nil || year < valorantRelease {
		year = 1970
	}
	return year
}

func fillEventDates(eventIds []*EventId) {
	for _, e := range eventIds {
		startYear := e.getYearFromTitle()
		start, end := e.getDayAndMonth()
		start = start.AddDate(startYear, 0, 0)
		end = end.AddDate(startYear, 0, 0)
		if start.Month() == time.December && end.Month() == time.January {
			end = end.AddDate(1, 0, 0)
		}
		e.Start = start
		e.End = end
	}

	withoutYears := []*EventId{}
	current := time.Time{}
	for _, e := range eventIds {
		if e.Start.Year() == 1970 {
			withoutYears = append(withoutYears, e)
		} else {
			current = e.Start
			setYearFromCurrent(withoutYears, current)
			withoutYears = []*EventId{}
		}
	}

	if len(withoutYears) > 0 {
		setYearFromCurrent(withoutYears, current)
	}
}

func setYearFromCurrent(eventIds []*EventId, current time.Time) {
	for _, e := range eventIds {
		if current.Month() == time.December && e.Start.Month() == time.January {
			e.Start = e.Start.AddDate(current.Year()-1969, 0, 0)
			e.End = e.End.AddDate(current.Year()-1969, 0, 0)
		} else if current.Month() == time.January && e.Start.Month() == time.December {
			e.Start = e.Start.AddDate(current.Year()-1971, 0, 0)
			e.End = e.End.AddDate(current.Year()-1971, 0, 0)
		} else {
			e.Start = e.Start.AddDate(current.Year()-1970, 0, 0)
			e.End = e.End.AddDate(current.Year()-1970, 0, 0)
		}
	}
}
