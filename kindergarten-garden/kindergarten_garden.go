package kindergarten

import (
	"errors"
	"slices"
)

// Garden type defines students and there plants.
type Garden map[string][]string

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
	newDiagram, err := sanitize(diagram)
	if err != nil {
		return nil, err
	}

	numOfPlants, numOfChildren := len(newDiagram), len(children)
	if numOfPlants/4 != numOfChildren {
		return nil, errInvalidDiagram
	}

	// make a new copy of children slice to prevent mutation of original slice.
	newChildren := make([]string, numOfChildren)
	copy(newChildren, children)

	slices.Sort(newChildren)
	ourGarden := make(Garden)

	for idx, child := range newChildren {
		if _, ok := ourGarden[child]; ok {
			return nil, errors.New("name was repeated in input")
		}
		plants := make([]string, 0, 4)
		firstRowPos := idx * 2
		secondRowPos := (numOfPlants / 2) + firstRowPos
		for _, val := range []int{firstRowPos, secondRowPos} {
			for j := val; j < val+2; j++ {
				plant, ok := plantLookup[newDiagram[j]]
				if !ok {
					return nil, errInvalidDiagram
				}
				plants = append(plants, plant)
			}
		}
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

// sanitize takes a diagram and sanitizes it by removing '\n', it  returns an error if diagram
// isn't proper.
func sanitize(diagram string) ([]byte, error) {
	newDiagram := []byte(diagram)
	rowLen := len(newDiagram) / 2

	if (rowLen-1)%2 != 0 || newDiagram[rowLen] != '\n' {
		return nil, errInvalidDiagram
	}
	newDiagram = append(newDiagram[1:rowLen], newDiagram[rowLen+1:]...)
	return newDiagram, nil
}
