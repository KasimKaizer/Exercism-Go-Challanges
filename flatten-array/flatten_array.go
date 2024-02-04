package flatten

func Flatten(nested interface{}) []interface{} {

	newArray := make([]interface{}, 0)

	for _, item := range nested.([]interface{}) {
		switch v := item.(type) {
		case []interface{}:
			newArray = append(newArray, Flatten(v)...)
		case nil:
			break
		default:
			newArray = append(newArray, v)
		}
	}
	return newArray
}
