// Package pascal contains solution for Pascal's Triangle exercise on Exercism.
package pascal

// Triangle generates a pascal's triangle which represents the input number.
func Triangle(n int) [][]int {
	if n < 1 {
		return nil // we can't generate a triangle for a number lower then 1.
	}
	output := make([][]int, n)
	output[0] = []int{1} // add the first row of the triangle.
	for i := 1; i < n; i++ {
		// generate a new row based on previous row and add it to output.
		output[i] = newRow(output[i-1])
	}
	return output
}

// newRow takes a row of pascal's triangle and generates it next row.
func newRow(prevArray []int) []int {
	prevLen := len(prevArray)
	output := make([]int, prevLen+1)  // each row of pascal's triangle is 1 bigger then previous row.
	output[0], output[prevLen] = 1, 1 // add '1' to start and end of the row.

	// we only need to iterate through half of previous row as each row in pascal's triangle is
	// symmetrical.
	for i := 1; i <= (prevLen)/2; i++ {
		// get the value by adding the two adjacent values in prev row.
		val := (prevArray[i-1] + prevArray[i])
		output[i], output[prevLen-i] = val, val
	}
	return output
}
