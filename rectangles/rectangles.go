// Package rectangles contains tools to count number of rectangles in a given diagram
package rectangles

// '+' represent a corner
//
// iterate through the diagram line by line, for each '+' we encounter, find it pairs in
// vertical axis, then from the points of vertical axis find its pairs, in horizontal axis.
// valid rectangle, horizontal = "-", "+", vertical =  "|", "+".
//
// functions:
// find potential points.
// find complementary vertical points.
// find complementary horizontal points.
// confirm if rectangle.

// position represents a point in diagram
type position struct {
	row int
	col int
}

// Count returns the number of rectangles in the provided diagram.
func Count(diagram []string) int {
	count := 0
	for i := 0; i < len(diagram)-1; i++ {
		for j := 0; j < len(diagram[i])-1; j++ {
			if diagram[i][j] != '+' {
				continue
			}
			count += countRect(diagram, position{i, j})
		}
	}
	return count
}

// countRect counts number of rectangles from a given point
func countRect(diagram []string, point position) int {
	count := 0
	for _, ver := range findVer(diagram, point) {
		for _, hoz := range findHoriz(diagram, ver) {
			if confirmRect(diagram, point, hoz) {
				count++
			}
		}
	}
	return count
}

// findVer finds all the complementary vertical points of a point. which could form a
// vertical side of a rectangle
func findVer(diagram []string, point position) []position {
	out := make([]position, 0)
	for i := point.col + 1; i < len(diagram[point.row]); i++ {
		if diagram[point.row][i] == '+' {
			out = append(out, position{point.row, i})
			continue
		}
		if diagram[point.row][i] != '-' {
			break
		}
	}
	return out
}

// findHoriz finds all the complementary horizontal points of a point which could form a
// horizontal side of a rectangle
func findHoriz(diagram []string, point position) []position {
	out := make([]position, 0)
	for i := point.row + 1; i < len(diagram); i++ {
		if diagram[i][point.col] == '+' {
			out = append(out, position{i, point.col})
			continue
		}
		if diagram[i][point.col] != '|' {
			break
		}
	}
	return out
}

// confirmRect confirms if the given points form a rectangle.
func confirmRect(diagram []string, main, hori position) bool {
	if diagram[hori.row][main.col] != '+' {
		return false
	}
	// walk backwards vertically
	for i := hori.col - 1; i > main.col; i-- {
		if diagram[hori.row][i] != '+' && diagram[hori.row][i] != '-' {
			return false
		}
	}
	// walk backwards horizontally
	for i := hori.row - 1; i > main.row; i-- {
		if diagram[i][main.col] != '+' && diagram[i][main.col] != '|' {
			return false
		}
	}
	return true
}
