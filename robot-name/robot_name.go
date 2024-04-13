// package robotname contains solution for Robot Name exercise on Exercism.
package robotname

import (
	"errors"
	"math/rand"
	"strings"
	"sync"
)

// limit defines the number of names possible.
const limit = 26 * 26 * 10 * 10 * 10

// memory stores all then names current in use by robots.
type memory struct {
	data map[string]struct{}
	mx   sync.Mutex
}

func (m *memory) exists(name string) bool {
	m.mx.Lock()
	defer m.mx.Unlock()
	_, ok := m.data[name]
	return ok
}

func (m *memory) set(name string) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.data[name] = struct{}{}
}

func (m *memory) len() int {
	m.mx.Lock()
	defer m.mx.Unlock()
	return len(m.data)
}

// Robot defines a robot, it has only one field which is its name.
type Robot struct {
	name string
}

var data = &memory{data: make(map[string]struct{})}

// Name method returns the name of the robot, if it doesn't have a name then it would
// generate a random unique name.
func (r *Robot) Name() (string, error) {

	if r.name != "" {
		return r.name, nil
	}

	if limit == data.len() {
		return "", errors.New("maximum limit of unique names that could be created reached")
	}
	r.name = r.nameGen()

	return r.name, nil
}

// Reset method reset's the robot to its default settings. i.e. it replaces the robot's current
// name with a random unique name.
func (r *Robot) Reset() error {
	// cause of the restrictions of this exercise, this would cause tests to fail. even tho its
	// sensible.
	// delete(memory, r.name)
	r.name = ""
	_, err := r.Name()
	if err != nil {
		return err
	}
	return nil
}

// nameGen generates a random unique name in a specific format and returns it.
func (r *Robot) nameGen() string {
	var output strings.Builder
	for i := 0; i < 2; i++ {
		output.WriteByte(byte('A') + byte(rand.Intn(26)))
	}
	for i := 0; i < 3; i++ {
		output.WriteByte(byte('0') + byte(rand.Intn(10)))
	}
	name := output.String()
	if data.exists(name) {
		return r.nameGen()
	}
	data.set(name)
	return name
}
