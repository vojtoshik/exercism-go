package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

// Tally takes as input results of the matches in the following format:
//
// Allegoric Alaskans;Blithering Badgers;win
// Devastating Donkeys;Courageous Californians;draw
// Devastating Donkeys;Allegoric Alaskans;win
// Courageous Californians;Blithering Badgers;loss
// Blithering Badgers;Devastating Donkeys;loss
// Allegoric Alaskans;Courageous Californians;win
//
// And creates a table with summaries of all games
func Tally(input io.Reader, writer io.Writer) error {

	gr := newGamesRecords()
	buffer, err := ioutil.ReadAll(input)

	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(buffer), "\n") {

		line = strings.Trim(line, " ")

		if line == "" || line[0] == '#' {
			continue
		}

		chunks := strings.Split(line, ";")

		if len(chunks) != 3 {
			return errors.New("record has unexpected format: " + line)
		}

		if !isGameResult(chunks[2]) {
			return errors.New("unknown game outcome: " + chunks[2])
		}

		gr.add(chunks[0], chunks[1], chunks[2])
	}

	gr.writeTo(writer)

	return nil
}

type teamRecord struct {
	teamName string
	games    map[string]int
}

func newTeamRecord(teamName string) *teamRecord {
	return &teamRecord{
		teamName: teamName,
		games: map[string]int{
			"win":  0,
			"loss": 0,
			"draw": 0,
		},
	}
}

func (s *teamRecord) points() int {
	return s.games["win"]*3 + s.games["draw"]
}

func (s *teamRecord) gamesPlayed() int {
	return s.games["win"] + s.games["draw"] + s.games["loss"]
}

type resultsTable map[string]*teamRecord

func newGamesRecords() resultsTable {
	return resultsTable{}
}

func (rt resultsTable) add(teamA, teamB, outcome string) {
	rt.addTeamRecord(teamA, outcome)
	rt.addTeamRecord(teamB, invertGameResult(outcome))
}

func (rt resultsTable) addTeamRecord(teamName, outcome string) {
	if _, ok := rt[teamName]; !ok {
		rt[teamName] = newTeamRecord(teamName)
	}

	teamRecord := rt[teamName]
	teamRecord.games[outcome]++
}

func (rt resultsTable) writeTo(writer io.Writer) {

	h := fmt.Sprintf("%-31s| MP |  W |  D |  L |  P\n", "Team")
	writer.Write([]byte(h))

	for _, v := range rt.getSortedResults() {

		tr := fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n",
			v.teamName,
			v.gamesPlayed(),
			v.games["win"],
			v.games["draw"],
			v.games["loss"],
			v.points())
		writer.Write([]byte(tr))
	}
}

func (rt resultsTable) getSortedResults() []teamRecord {

	var r []teamRecord

	for _, record := range rt {
		r = append(r, *record)
	}

	sort.Slice(r, func(i, j int) bool {

		if r[i].points() == r[j].points() {
			return r[i].teamName < r[j].teamName
		}

		return r[i].points() > r[j].points()
	})

	return r
}

func isGameResult(s string) bool {
	return s == "win" || s == "loss" || s == "draw"
}

func invertGameResult(gr string) string {
	switch gr {
	case "win":
		return "loss"
	case "loss":
		return "win"
	default:
		return "draw"
	}
}
