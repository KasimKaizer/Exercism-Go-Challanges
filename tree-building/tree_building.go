// Package tree contains solution for the Tree Building exercise on Exercism.
package tree

import (
	"errors"
	"sort"
)

var (
	// error for badly improper sequenced records.
	errNonSeq = errors.New("records are not sequential")
)

// Record holds the its own id and its parent's id.
type Record struct {
	ID     int
	Parent int
}

// Node holds its own ID and all its children in form of a sorted array.
type Node struct {
	ID       int
	Children []*Node
}

// Build takes a list of records and turns it into a tree of nodes.
func Build(records []Record) (*Node, error) {
	recLen := len(records)
	if recLen == 0 {
		return nil, nil // return nil if record list is empty.
	}
	recordsCopy := make([]Record, recLen)
	newRecords := make([]*Node, recLen)
	// make a copy of the record slice, so we don't mutate the original record slice when we sort records.
	copy(recordsCopy, records)

	// sort the copy of records.
	sort.Slice(recordsCopy, func(i, j int) bool {
		return recordsCopy[i].ID < recordsCopy[j].ID
	})

	// range over all the sorted records.
	for idx, record := range recordsCopy {
		// if index of the record is not equal to its id then we know the records are not
		// properly sequential.
		if idx != record.ID ||
			// records ID can't be less then its parents ID.
			(record.ID <= record.Parent && record.ID != 0) ||
			// first record '0' can't have any parent other then itself.
			(record.ID == 0 && record.Parent != 0) {
			return nil, errNonSeq
		}
		// add current record as a node to the newRecord slice at the current index
		newRecords[idx] = &Node{ID: record.ID}
		if idx == 0 {
			// if current record is '0' then it can't be a child of any other record so continue to next iteration.
			continue
		}
		// get the parent node from the newRecords slice using the parents id in current record,
		// then append the current node as a child of that parent.
		// we don't have to worry about errors here as this would only trigger if parents id is
		// less then current records id.
		newRecords[record.Parent].Children = append(newRecords[record.Parent].Children, newRecords[record.ID])
	}

	return newRecords[0], nil // return the first node from the newRecords slice.
}
