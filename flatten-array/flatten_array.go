// Package flatten contains tools to flatten an array containing any type.
package flatten

// Flatten takes a nested array and flattens it into a single array with no nesting.
func Flatten(nested any) []any {
	if _, ok := nested.([]any); !ok {
		return []any{nested}
	}

	output := make([]any, 0)

	for _, item := range nested.([]any) {
		switch v := item.(type) {
		case []any:
			output = append(output, Flatten(v)...)
		case nil:
			break
		default:
			output = append(output, v)
		}
	}
	return output
}
