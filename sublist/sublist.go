// Package sublist contains solution for Sub List exercise on Exercism.
package sublist

// My though process to solve this. I am leaving this in for future me.
// BREAK DOWN THE PROBLEM INTO SUB PROBLEMS.
// check if both lists are of equal length, if yes check if they are equal
// if they aren't of equal length find the biggest string
// create a loop to find the starting point of the small string in big string
// return based on which string was bigger
// if lists were unsorted, we would sort them and do this same process.

// Sublist takes two lists and tells you about the relation between them.
func Sublist(l1, l2 []int) Relation {

	l1Len := len(l1)
	l2Len := len(l2)

	switch {

	case l1Len == l2Len:
		if subset(l1, l2) { // if both list are of same length
			return RelationEqual // return Equal is they are equal, else just terminate switch.
		}
	case l1Len > l2Len:
		if subset(l1, l2) { // if l1 is bigger then l2, check if l2 is part of l1.
			return RelationSuperlist // return SuperList if it is.
		}
	case l1Len < l2Len:
		if subset(l2, l1) { // if l2 is bigger then l1, check if l1 is part of l2.
			return RelationSublist // return SubList if it is.
		}
	}

	return RelationUnequal // return unequal if non of the above cases return.

}

// subset takes 2 sets and returns if smallerList is subset of largerList.
func subset(largerList, smallerList []int) bool {

	if len(smallerList) == 0 {
		return true // if smallerList is nil then it is a subset of larger.
	}
	largerLen := len(largerList)
	smallerLen := len(smallerList)
	/*
		if smallerLen > largerLen {
			return false // if largerList is smaller then smallerList, then smallerList can't be a subset.
		}
	*/
	// we iterate through all possible subsets in largerList which satisfy the length of smallerList
	for i := 0; i < (largerLen - smallerLen + 1); i++ {
		isSubset := true // we assume it is a subset.

		for idx, item := range smallerList {
			if largerList[idx+i] != item { // if the num of current subset of largerList is not in smallerList
				isSubset = false // we know our assumption was false
				break            // we continue to next iteration of subset.
			}
		}

		if isSubset { // if in above loop if condition was never triggered then we know it is a subset.
			return true // we return true
		}
	}
	// if we went trough all possible iteration and 'isSubset' if condition never trigger then we know it's not a
	// subset.
	return false
}
