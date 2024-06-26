package api

import (
	"strings"

	proto "github.com/derarken/vlr-api/gen/vlr/api"
	"github.com/derarken/vlr-api/src/scrapers"
)

func GetTeam(teamId string) (*proto.Team, error) {
	scrapedTeam, err := scrapers.ScrapeTeam(teamId)
	if err != nil {
		return nil, err
	}

	team := &proto.Team{
		Head: &proto.Team_Head{
			TeamId:  teamId,
			Name:    scrapedTeam.Name,
			Tricode: scrapedTeam.Tricode,
			Region:  scrapedTeam.Region,
		},
		Roster: &proto.Team_Roster{
			Players: getRosterEntries(scrapedTeam.Players),
			Staff:   getRosterEntries(scrapedTeam.Staff),
		},
	}

	return team, nil
}

func getRosterEntries(entries []*scrapers.RosterEntry) []*proto.Team_Roster_Entry {
	rosterEntries := make([]*proto.Team_Roster_Entry, 0, len(entries))

	for _, player := range entries {
		rosterEntries = append(rosterEntries, &proto.Team_Roster_Entry{
			PlayerId: player.PlayerID,
			Name:     player.PlayerName,
			RealName: player.RealName,
			Role:     getRole(player.Role),
		})
	}

	return rosterEntries
}

func getRole(role string) proto.RosterEntryRole {
	switch strings.ToLower(role) {
	case "inactive":
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_INACTIVE
	case "":
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_PLAYER
	case "sub":
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_SUB
	case "stand-In":
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_STAND_IN
	case "coach":
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_COACH
	case "head coach":
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_HEAD_COACH
	case "assistant coach":
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_ASSISTANT_COACH
	case "performance coach":
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_PERFORMANCE_COACH
	case "analyst":
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_ANALYST
	case "manager":
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_MANAGER
	default:
		return proto.RosterEntryRole_ROSTER_ENTRY_ROLE_UNSPECIFIED
	}
}
