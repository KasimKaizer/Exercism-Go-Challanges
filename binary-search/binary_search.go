package binarysearch

// notFound contains result for in case a search returns no result.
const notFound = -1

// SearchInts searches the given key in the given array and returns it position in the array.
// it uses binary search to search the array.
func SearchInts(list []int, key int) int {
	start, end := 0, len(list)-1
	for end >= 0 {
		middle := (end + start) / 2

		if list[middle] == key {
			return middle
		}

		if start == end {
			return notFound
		}

		if list[middle] > key {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return notFound
}
