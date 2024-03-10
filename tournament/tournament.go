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
	MP   int // match played
	W    int // wins
	D    int // draws
	P    int // points
}

// teams type hold on all the teams taking part in tournament.
type teams map[string]*team

// newTeams creates a new teams object.
func newTeams() teams {
	teams := make(teams)
	teamNames := []string{"Devastating Donkeys", "Allegoric Alaskians", "Blithering Badgers", "Courageous Californians"} // list of all teams.

	for _, name := range teamNames {
		teams[name] = &team{Name: name}
	}
	return teams
}

// win method executes the win condition on a team.
func (t *team) win() {
	t.MP++
	t.W++
	t.P += 3
}

// draw method executes the draw condition on a team.
func (t *team) draw() {
	t.MP++
	t.D++
	t.P++
}

// lose method executes the lose condition on a team.
func (t *team) lose() {
	t.MP++
}

// calLose method calculates the amount of loses of a team.
func (t *team) calLose() int {
	return t.MP - (t.W + t.D)
}

var (
	eerInvalidData = errors.New("tournament: invalid data submitted, cannot parse")
)

// Tally converts the data of matches into a formatted table.
func Tally(reader io.Reader, writer io.Writer) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	matches := strings.Split(buf.String(), "\n")
	teams := newTeams()
	for _, match := range matches {
		err := teams.updateScore(match)
		if err != nil {
			return err
		}
	}
	return presentData(teams.sortData(), writer)

}

// updateScore take the data for a single match and teams object and update the teams object based
// on the match data.
func (t *teams) updateScore(match string) error {

	// strings starting with '#' are counted as comments.
	if match == "" || strings.HasPrefix(match, "#") {
		return nil
	}
	data := strings.Split(match, ";")

	// the first two elements would be team names
	winner, ok := (*t)[data[0]]
	if !ok {
		return eerInvalidData
	}
	loser, ok := (*t)[data[1]]
	if !ok {
		return eerInvalidData
	}

	// last element would be the outcome
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
func (t *teams) sortData() []*team {
	teamSlice := make([]*team, 0, len(*t))
	for _, value := range *t {
		teamSlice = append(teamSlice, value)
	}

	// sort the slice based on the each team's points, its in descending order.
	sort.Slice(teamSlice, func(i, j int) bool {
		if teamSlice[i].P == teamSlice[j].P {
			return teamSlice[i].Name < teamSlice[j].Name
		}
		return teamSlice[i].P > teamSlice[j].P
	})
	return teamSlice
}

// presentData formats the provided teams data and writes it to the writer.
func presentData(teams []*team, writer io.Writer) error {
	var output strings.Builder
	output.WriteString(
		"Team                           | MP |  W |  D |  L |  P\n", // this is the header.
	)
	for _, team := range teams { // for each team in slice, we write out data into string.
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
	_, err := writer.Write([]byte(output.String()))
	return err
}
