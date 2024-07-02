package api

import (
	"errors"
	"fmt"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
)

func GetEventIds(request *proto.GetEventIdsRequest) ([]string, error) {
	switch request.State {
	case proto.EventState_EVENT_STATE_ONGOING, proto.EventState_EVENT_STATE_UPCOMING:
		return getUpcomingOrOngoingEventIds()

	case proto.EventState_EVENT_STATE_COMPLETED:
		return getCompletedEventIds()

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

func getCompletedEventIds() ([]string, error) {
	return nil, errors.New("not implemented")
}

func getUpcomingOrOngoingEventIds() ([]string, error) {
	return nil, errors.New("not implemented")
}
