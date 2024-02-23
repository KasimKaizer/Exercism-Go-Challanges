// package robotname contains solution for Robot Name exercise on Exercism.
package robotname

import (
	"errors"
	"math/rand"
	"strings"
)

// limit defines the number of names possible.
const limit = 26 * 26 * 10 * 10 * 10

// memory stores all then names current in use by robots.
var memory = make(map[string]bool)

// Robot defines a robot, it has only one field which is its name.
type Robot struct {
	name string
}

// Name method returns the name of the robot, if it doesn't have a name then it would
// generate a random unique name.
func (r *Robot) Name() (string, error) {

	if r.name != "" {
		return r.name, nil // check if the robot already has a name, if yes then return it.
	}

	if limit == len(memory) { // return error if there aren't any names which could be generated.
		return "", errors.New("maximum limit of unique names that could be created reached")
	}
	r.name = r.nameGen() // get an unique name, and set it as the robot's name

	return r.name, nil
}

// Reset method reset's the robot to its default settings. i.e. it replaces the robot's current
// name with a random unique name.
func (r *Robot) Reset() error {
	// cause of the restrictions of this exercise, this would cause tests to fail. even tho its
	// sensible.
	// delete(memory, r.name) // free the current name for other robot's to use.
	r.name = ""        // remove the current robot's name.
	_, err := r.Name() // give a new unique name to the bot.
	if err != nil {
		return err
	}
	return nil
}

// nameGen generates a random unique name in a specific format and returns it.
func (r *Robot) nameGen() string {

	var output strings.Builder
	for i := 0; i < 2; i++ { // add two random capital letters to the string.
		output.WriteRune('A' + rune(rand.Intn(26)))
	}
	for i := 0; i < 3; i++ { // add three random numbers to the string.
		output.WriteRune('0' + rune(rand.Intn(10)))
	}
	name := output.String()
	if memory[name] {
		return r.nameGen() // keep on generating until we get a unique name.
	}
	memory[name] = true // add the name to the memory.
	return name
}
