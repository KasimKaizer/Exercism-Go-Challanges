package kindergarten

import (
	"errors"
	"slices"
	"strings"
)

// Garden type defines students and there plants.
type Garden map[string][]string

// error for an invalid diagram.
var errInvalidDiagram = errors.New("invalid diagram provided")

// plantLookup maps bytes which contains the initial letter for a plant to there full name.
var plantLookup = map[byte]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

// NewGarden takes the diagram of plants arrangement and a list of student names
// and returns a map mapping a student with the plants they should be assigned.
func NewGarden(diagram string, children []string) (*Garden, error) {
	newDiagram, err := sanitize(diagram) // sanitize the input diagram so it can be parsed.
	if err != nil {
		return nil, err // elevate the error if its encountered.
	}
	// if the number of grouped plants in a single row doesn't equal the number of students then
	// the diagram is invalid.
	if len(newDiagram[0]) != len(children) {
		return nil, errInvalidDiagram
	}
	// we make a new copy of children slice to prevent mutation of original slice.
	newChildren := make([]string, len(children))
	copy(newChildren, children)
	// sort the new slice alphabetically.
	slices.Sort(newChildren)
	ourGarden := make(Garden)
	// we iterate through all the students. idx of a student in students list would align with
	// the idx of there plants in both rows of the newDiagram.
	for idx, child := range newChildren {
		plants := make([]string, 0)
		if _, ok := ourGarden[child]; ok {
			// we make sure any child was not repeated, if it was then the students list is
			// invalid.
			return nil, errors.New("name was repeated in input")
		}
		for _, row := range newDiagram { // we iterate here twice, for both rows
			// as idx of child is equal to idx of there plants in the row.
			for _, bt := range row[idx] {
				// we will iterate here twice as well for both plants that belong to a child in a
				// single row.
				plant, ok := plantLookup[bt] // look up the name of that plant
				if !ok {
					// if it doesn't exist in plant map then diagram is invalid
					return nil, errInvalidDiagram
				}
				// TODO: potential optimization, use an array of size 4.
				plants = append(plants, plant) // append that plant to plants slice.
			}
		}
		// add the current plant slice with the key as child's name to the map .
		ourGarden[child] = plants
	}
	return &ourGarden, nil
}

// Plants method takes the name of the child as input and returns the plants they are
// responsible for.
func (g Garden) Plants(child string) ([]string, bool) {
	plants, ok := g[child]
	return plants, ok
}

// splitBy2 takes a slice of bytes indicating a single row of plans and returns a matrix
// which groups two plants together.
func splitBy2(str []byte) [][]byte {
	output := make([][]byte, 0)
	for i := 0; i < len(str); i = i + 2 {
		// we iterate over all the bytes & append a slice grouping 2 plants in one slice
		// to a matrix.
		output = append(output, str[i:i+2])
	}
	return output
}

// sanitize takes the initial diagram and sanitizes it so it can
// parsed by our program. it does this by
// First: split the diagram up in two rows.
// Second: group two plants together on each row, indicating students plants.
// Third: combining all of that into a single 3D matrix and returning it.
func sanitize(diagram string) ([][][]byte, error) {
	rows := strings.Split(diagram, "\n") // split the diagram by newline.
	// if the number rows we got is less then 3 then the diagram is invalid
	if len(rows) < 3 {
		return nil, errInvalidDiagram
	}
	rows = rows[1:] // first row would be a empty string, we get rid of that.
	if len(rows[0]) != len(rows[1]) {
		// if the length of both rows are not equal then diagram is invalid.
		return nil, errInvalidDiagram
	}
	// send both rows to splitBy2 func. here each split indicates a single students
	// plants on that row.
	row1Slice := splitBy2([]byte(rows[0]))
	row2Slice := splitBy2([]byte(rows[1]))

	// return the 3D matrix, where first slice indicates the number of rows, second slice indicates
	// the plants of a all student grouped on that perticular row and the third slice indicates
	// the plants of a single student on that row.
	return [][][]byte{row1Slice, row2Slice}, nil
}
