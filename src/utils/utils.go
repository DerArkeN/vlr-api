package utils

import (
	"strconv"
	"strings"
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
