package main

import (
	"fmt"
	"time"

	"github.com/derarken/vlr-api/proto"
	"github.com/derarken/vlr-api/src/api"
)

func main() {
	ids, err := api.GetMatchIds(proto.Status_STATUS_UPCOMING, time.Now().Add(time.Hour*24*30), time.Now().Add(time.Hour*24*60))
	if err != nil {
		panic(err)
	}
	fmt.Println(ids)
}

// matches, err := scraper_matches.ScrapeEventMatches("2004/champions-tour-2024-americas-stage-1")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(matches)

// 	mapsWonWithBonusRound := 0
// 	mapsLostWithBonusRound := 0

// 	for i, match := range matches {
// 		detail, err := scraper_match.ScrapeMatchDetail(match.Id)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		for _, map_ := range detail.Maps {
// 			if len(map_.Rounds) == 0 {
// 				continue
// 			}

// 			bonusWonTeam1 := false
// 			bonusWonTeam2 := false
// 			for _, round := range map_.Rounds {
// 				if round == "" {
// 					continue
// 				}

// 				if round == "3-0" {
// 					bonusWonTeam1 = true
// 				}

// 				if round == "0-3" {
// 					bonusWonTeam2 = true
// 				}
// 			}

// 			lastRound := map_.Rounds[len(map_.Rounds)-1]
// 			score1, score2, err := utils.GetScoreForRound(lastRound)
// 			if err != nil {
// 				log.Println(err)
// 				continue
// 			}

// 			if bonusWonTeam1 {
// 				if score1 > score2 {
// 					mapsWonWithBonusRound += 1
// 				} else {
// 					mapsLostWithBonusRound += 1
// 				}
// 			}

// 			if bonusWonTeam2 {
// 				if score2 > score1 {
// 					mapsWonWithBonusRound += 1
// 				} else {
// 					mapsLostWithBonusRound += 1
// 				}
// 			}
// 		}
// 		log.Println(strconv.Itoa(i+1) + "/" + strconv.Itoa(len(matches)))
// 	}

// 	log.Println("Maps won with bonus round:", mapsWonWithBonusRound)
// 	log.Println("Maps lost with bonus round:", mapsLostWithBonusRound)
// 	log.Println("Winrate with bonus round:", float64(mapsWonWithBonusRound)/float64(mapsWonWithBonusRound+mapsLostWithBonusRound)*100, "%")
