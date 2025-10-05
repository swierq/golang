package premierleague

import (
	"fmt"
	"net/http"
	"time"
)

type Fixtures []struct {
	Code                 int       `json:"code"`
	Event                int       `json:"event"`
	Finished             bool      `json:"finished"`
	FinishedProvisional  bool      `json:"finished_provisional"`
	ID                   int       `json:"id"`
	KickoffTime          time.Time `json:"kickoff_time"`
	Minutes              int       `json:"minutes"`
	ProvisionalStartTime bool      `json:"provisional_start_time"`
	Started              bool      `json:"started"`
	AwayTeam             int       `json:"team_a"`
	AwayTeamScore        int       `json:"team_a_score"`
	HomeTeam             int       `json:"team_h"`
	HomeTeamScore        int       `json:"team_h_score"`
	Stats                []Stat    `json:"stats"`
	HomeTeamDifficulty   int       `json:"team_h_difficulty"`
	AwayTeamDifficulty   int       `json:"team_a_difficulty"`
	PulseID              int       `json:"pulse_id"`
}

type Stat struct {
	Identifier string        `json:"identifier"`
	A          []StatElement `json:"a"`
	H          []StatElement `json:"h"`
}

type StatElement struct {
	Value   int `json:"value"`
	Element int `json:"element"`
}

func AllFixtures(team string, httpClient *http.Client, daysBack, daysForward int) error {
	client := NewClient(httpClient)
	fixtures, err := client.GetFixtures()
	if err != nil {
		panic(err)
	}

	bootstrapData, err := client.GetBootstrapData()
	bootstrapData.TeamIdToName = make(map[int]string)
	bootstrapData.TeamNameToId = make(map[string]int)
	if err != nil {
		return err
	}

	if err != nil {
		panic(err)
	}

	now := time.Now()
	durationStart, _ := time.ParseDuration(fmt.Sprintf("-%dh", daysBack*24))
	durationEnd, _ := time.ParseDuration(fmt.Sprintf("%dh", daysForward*24))

	start := now.Add(durationStart)
	end := now.Add(durationEnd)

	fmt.Println("All Fixtures:")
	for _, fix := range fixtures {
		if !fix.KickoffTime.After(start) || !fix.KickoffTime.Before(end) {
			continue
		}
		home, err := bootstrapData.GetTeamName(fix.HomeTeam)
		if err != nil {
			return err
		}

		away, err := bootstrapData.GetTeamName(fix.AwayTeam)
		if err != nil {
			return err
		}
		if team != "" && home != team && away != team {
			continue
		}
		fmt.Printf("%s  %d : %d %s - %s", home, fix.HomeTeamScore, fix.AwayTeamScore, away, fix.KickoffTime)
		if fix.Finished {
			fmt.Println(" - FINISHED")
		} else if fix.Started {
			fmt.Println(" - PLAYING")
		} else {
			fmt.Println("")
		}
	}

	return nil
}
