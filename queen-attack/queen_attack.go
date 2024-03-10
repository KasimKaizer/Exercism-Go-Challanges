// Package queenattack contains solution for Queen Attack exercise on Exercism.
package queenattack

import (
	"errors"
	"math"
)

// position represents a position of a chess piece on a board.
type position struct {
	col int
	row int
}

var errInvPos = errors.New("queenattack: invalid position for a chess piece provided")

// CanQueenAttack takes position of two queens on a chess board and returns true or false based on
// if those queen can attack each other.
func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition {
		return false, errInvPos
	}
	wPos, err := newPosition(whitePosition)
	if err != nil {
		return false, err
	}
	bPos, err := newPosition(blackPosition)
	if err != nil {
		return false, err
	}
	return wPos.col == bPos.col || wPos.row == bPos.row || wPos.isDiagonal(bPos), nil
}

// newPosition takes a position of a chess piece represented in a conventional way and converts it
// into a position object.
func newPosition(pos string) (*position, error) {
	if len(pos) != 2 || pos[0] > 'h' || pos[0] < 'a' || pos[1] < '1' || pos[1] > '8' {
		return nil, errInvPos
	}

	var output position
	output.col = int(pos[0]-'a') + 1
	output.row = int(pos[1] - '0')

	return &output, nil
}

// isDiagonal methods returns true if the position passed is diagonal to the position
// this method is called on.
func (p *position) isDiagonal(pos *position) bool {
	return math.Abs(float64(p.col)-float64(p.row)) == math.Abs(float64(pos.col)-float64(pos.row))
}
