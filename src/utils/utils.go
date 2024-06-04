package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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

func AbbreviationToLocation(abbreviation string) (*time.Location, error) {
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
