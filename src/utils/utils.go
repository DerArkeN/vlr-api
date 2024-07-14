package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrFromAfterTo = errors.New("from time must not be after to time")
	ErrToInFuture  = errors.New("to time must not be in the future when state is completed")
)

func PrettifyString(input string) string {
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")
	input = strings.TrimSpace(input)
	return input
}

func GetScoreForRound(round string) (int, int, error) {
	scores := strings.Split(round, "-")

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

func ShouldValidateTime(loc *time.Location, from time.Time, to time.Time) bool {
	return loc != nil && !from.IsZero() && !to.IsZero()
}

func IsTimeInRange(from time.Time, to time.Time, time time.Time) bool {
	return time.UnixMilli() >= from.UnixMilli() && time.UnixMilli() <= to.UnixMilli()
}

func ValidateCompletedTimes(frompb *timestamppb.Timestamp, topb *timestamppb.Timestamp, duration time.Duration) (time.Time, time.Time, error) {
	from := time.Time{}
	if frompb != nil {
		from = frompb.AsTime()
	}
	to := time.Time{}
	if topb != nil {
		to = topb.AsTime()
	}

	if from.IsZero() && to.IsZero() {
		to = time.Now()
		from = to.Add(duration)
	}
	if from.IsZero() && !to.IsZero() {
		from = to.Add(duration)
	}
	if !from.IsZero() && to.IsZero() {
		to = time.Now()
	}
	if from.After(to) {
		return time.Time{}, time.Time{}, ErrFromAfterTo
	}
	if to.After(time.Now()) {
		return time.Time{}, time.Time{}, ErrToInFuture
	}
	return from, to, nil
}

func ValidateUpcomingTimes(frompb *timestamppb.Timestamp, topb *timestamppb.Timestamp, duration time.Duration) (time.Time, time.Time, error) {
	from := time.Time{}
	if frompb != nil {
		from = frompb.AsTime()
	}
	to := time.Time{}
	if topb != nil {
		to = topb.AsTime()
	}

	if from.IsZero() && to.IsZero() {
		from = time.Now()
		to = from.Add(duration)
	}
	if from.IsZero() && !to.IsZero() {
		from = time.Now()
	}
	if !from.IsZero() && to.IsZero() {
		to = from.Add(duration)
	}
	if from.After(to) {
		return time.Time{}, time.Time{}, ErrFromAfterTo
	}
	return from, to, nil
}
