// Package tournament contains solution for Tournament exercise on Exercism.
package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// team types defines a team and there stats.
type team struct {
	Name string
	MP   int
	W    int
	D    int
	P    int
}

// teams type hold on all the teams taking part in tournament.
type teams map[string]*team

// win method executes the win condition on a team, it increments the teams wins count by 1
// and adds 3 to its total points. it also increments the team's match played count  by 1.
func (t *team) win() {
	t.MP++
	t.W++
	t.P += 3
}

// draw method executes the draw condition on a team, it increments the teams draw count by 1
// and adds 1 to its total points. it also increments the team's match played count  by 1.
func (t *team) draw() {
	t.MP++
	t.D++
	t.P++
}

// lose method executes the lose condition on a team, it just increments the teams
// match played count and nothing else.
func (t *team) lose() {
	t.MP++
}

// calLose method calculates the amount of loses of a team.
func (t *team) calLose() int {
	return t.MP - (t.W + t.D)
}

var (
	// eerInvalidData denoted that invalid data was passed to the program.
	eerInvalidData = errors.New("invalid data submitted, cannot parse")
	// list of all teams. if in future, we want to increase the amount of teams
	// that are taking part in this tournament, then we would just add there names to this slice.
	teamNames = []string{"Devastating Donkeys", "Allegoric Alaskians", "Blithering Badgers", "Courageous Californians"}
)

// Tally converts the data of matches into a formatted table.
func Tally(reader io.Reader, writer io.Writer) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)                         // create a new buffer and read into that buffer.
	matches := strings.Split(buf.String(), "\n") // split the data by new line.

	// create a new teams object with all the teams taking part in tournament.
	teams := newTeams()

	for _, match := range matches { // go through each match in data
		err := updateScore(match, teams) // update the teams object based on that data.
		if err != nil {
			return err
		}
	}

	sortedTeams := sortData(teams) // sort our teams object and convert it into an ordered slice.
	var output strings.Builder
	output.WriteString(
		"Team                           | MP |  W |  D |  L |  P\n",
	) // this is our header.
	for _, team := range sortedTeams { // for each team in our slice, we write out data into string.
		output.WriteString(fmt.Sprintf(
			"%-31s| %2d | %2d | %2d | %2d | %2d\n",
			team.Name,
			team.MP,
			team.W,
			team.D,
			team.calLose(),
			team.P,
		))
	}
	writer.Write([]byte(output.String())) // we pass our data to writer.
	return nil
}

// updateScore take the data for a single match and teams object and update the teams object based
// on the match data.
func updateScore(match string, teams teams) error {
	// return nil for empty strings and any string starting with '#'.
	// strings starting with '#' are counted as comments.
	if match == "" || strings.HasPrefix(match, "#") {
		return nil
	}
	data := strings.Split(match, ";") // split our data by ';'
	// the first two items would be team names, check if they are valid names, if not
	// then return an error.
	winner, ok := teams[data[0]]
	if !ok {
		return eerInvalidData
	}

	loser, ok := teams[data[1]]
	if !ok {
		return eerInvalidData
	}

	// last element would be the outcome, preform the win, lose or draw condition on both
	// teams based on the outcome.
	switch data[2] {
	case "win":
		break
	case "loss":
		winner, loser = loser, winner
	case "draw":
		winner.draw()
		loser.draw()
		return nil
	default:
		return eerInvalidData
	}
	winner.win()
	loser.lose()
	return nil
}

// sortData takes an unordered and unsorted teams object and sorts it, and returns a ordered
// slice.
func sortData(teams teams) []*team {
	teamSlice := make([]*team, 0, len(teams))
	// copy all the elements from our teams object into a slice.
	for _, value := range teams {
		teamSlice = append(teamSlice, value)
	}

	// sort the slice based on the each team's points, where its ordered from higher to lower.
	sort.Slice(teamSlice, func(i, j int) bool {
		if teamSlice[i].P == teamSlice[j].P {
			return teamSlice[i].Name < teamSlice[j].Name
		}

		return teamSlice[i].P > teamSlice[j].P
	})
	return teamSlice
}

// newTeams creates a new teams object.
func newTeams() teams {
	teams := make(teams)
	for _, name := range teamNames {
		teams[name] = &team{Name: name}
	}
	return teams
}
