package api

import (
	"fmt"
	"time"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
	"github.com/derarken/vlr-api/src/scrapers"
	"github.com/derarken/vlr-api/src/utils"
)

type EventIdFactory struct{}

var eventIdFactory = &EventIdFactory{}

const eventIdsDefaultDuration = 7 * time.Hour * 24

func GetEventIds(request *proto.GetEventIdsRequest) ([]string, error) {
	switch request.State {
	case proto.EventState_EVENT_STATE_ONGOING:
		return eventIdFactory.getOngoingEventIds()

	case proto.EventState_EVENT_STATE_UPCOMING:
		from, to, err := utils.ValidateUpcomingTimes(request.From, request.To, eventIdsDefaultDuration)
		if err != nil {
			return nil, err
		}
		return eventIdFactory.getUpcomingEventIds(from, to)

	case proto.EventState_EVENT_STATE_COMPLETED:
		from, to, err := utils.ValidateCompletedTimes(request.From, request.To, eventIdsDefaultDuration)
		if err != nil {
			return nil, err
		}
		return eventIdFactory.getCompletedEventIds(from, to)

	default:
		validStates := []string{}
		for k, v := range proto.EventState_name {
			if k == 0 {
				continue
			}
			validStates = append(validStates, v)
		}
		return nil, fmt.Errorf("invalid state, valid states are %v", validStates)
	}
}

func (f *EventIdFactory) getOngoingEventIds() ([]string, error) {
	var eventIds []*scrapers.EventId
	page := 1
	for {
		var newEventIds []*scrapers.EventId
		var err error
		newEventIds, err = scrapers.ScrapeUpcomingOrOngoingEventIds(page)
		if err == scrapers.ErrNoEventIds {
			break
		}
		if err != nil {
			return nil, err
		}
		eventIds = append(eventIds, newEventIds...)

		lastEventId := newEventIds[len(newEventIds)-1]
		if lastEventId.State != string(scrapers.EVENT_STATE_ONGOING) {
			break
		}

		page++
	}

	return eventIdFactory.getEventIdsByStateAndTime(eventIds, scrapers.EVENT_STATE_ONGOING, nil, time.Time{}, time.Time{})
}

func (f *EventIdFactory) getUpcomingEventIds(from time.Time, to time.Time) ([]string, error) {
	loc, err := scrapers.GetLocation()
	if err != nil {
		return nil, err
	}

	var eventIds []*scrapers.EventId
	page := 1
	for {
		var newEventIds []*scrapers.EventId
		var err error
		newEventIds, err = scrapers.ScrapeUpcomingOrOngoingEventIds(page)
		if err == scrapers.ErrNoEventIds {
			break
		}
		if err != nil {
			return nil, err
		}
		eventIds = append(eventIds, newEventIds...)

		lastEventId := newEventIds[len(newEventIds)-1]
		start, _, err := lastEventId.GetStartAndEndUtc(loc)
		if err != nil {
			return nil, err
		}

		if start.After(to) {
			break
		}

		page++
	}

	return f.getEventIdsByStateAndTime(eventIds, scrapers.EVENT_STATE_UPCOMING, loc, from, to)
}

func (f *EventIdFactory) getCompletedEventIds(from time.Time, to time.Time) ([]string, error) {
	loc, err := scrapers.GetLocation()
	if err != nil {
		return nil, err
	}

	var eventIds []*scrapers.EventId
	page := 1
	for {
		var newEventIds []*scrapers.EventId
		var err error
		newEventIds, err = scrapers.ScrapeUpcomingOrOngoingEventIds(page)
		if err == scrapers.ErrNoEventIds {
			break
		}
		if err != nil {
			return nil, err
		}
		eventIds = append(eventIds, newEventIds...)

		lastEventId := newEventIds[len(newEventIds)-1]
		start, _, err := lastEventId.GetStartAndEndUtc(loc)
		if err != nil {
			return nil, err
		}

		if start.Before(to) {
			break
		}

		page++
	}

	return f.getEventIdsByStateAndTime(eventIds, scrapers.EVENT_STATE_UPCOMING, loc, from, to)
}

func (f *EventIdFactory) getEventIdsByStateAndTime(eventIds []*scrapers.EventId, state scrapers.EventState, loc *time.Location, from time.Time, to time.Time) ([]string, error) {
	var ids []string

	for _, eventId := range eventIds {
		if eventId.State != string(state) {
			continue
		}

		if utils.ShouldValidateTime(loc, from, to) {
			start, _, err := eventId.GetStartAndEndUtc(loc)
			if err != nil {
				return nil, err
			}

			if utils.IsTimeInRange(from, to, start) {
				ids = append(ids, eventId.EventId)
			}
		} else {
			ids = append(ids, eventId.EventId)
		}
	}

	return ids, nil
}
