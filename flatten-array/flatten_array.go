// Package flatten contains solution for Flatten Array Exercise on Exercism.
package flatten

// Flatten takes a nested array and flattens it into a single array with no nesting.
func Flatten(nested interface{}) []interface{} {
	if _, ok := nested.([]interface{}); !ok {
		return []interface{}{nested} // handle edge case, if input is not an array.
	}

	output := make([]interface{}, 0)

	for _, item := range nested.([]interface{}) { // loop through all items in the interface array.
		switch v := item.(type) { // type switch to check type of the current value.
		case []interface{}:
			// if the current item is an array then call Flatten() func
			// on it and append the returned array at the end of output array.
			output = append(output, Flatten(v)...)
		case nil:
			break // if the item if nil then move to next item.
		default:
			// if its non of the above then append that item to output array.
			output = append(output, v)
		}
	}
	return output
}
