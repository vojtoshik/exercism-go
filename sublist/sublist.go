package sublist

type Relation string

func Sublist(l1, l2 []int) Relation {
	if isEqual(l1, l2) {
		return "equal"
	} else if isSublist(l1, l2) {
		return "sublist"
	} else if isSublist(l2, l1) {
		return "superlist"
	}

	return "unequal"
}

func isEqual(l1, l2 []int) bool {

	if len(l1) != len(l2) {
		return false
	}

	for i, n1 := range l1 {
		if n1 != l2[i] {
			return false
		}
	}

	return true
}

func isSublist(l1, l2 []int) bool {

	if len(l1) == 0 {
		return true
	}

	if len(l1) > len(l2) {
		return false
	}

	for i := 0; i <= len(l2)-len(l1); i++ {
		if l2[i] == l1[0] && isEqual(l1, l2[i:i+len(l1)]) {
			return true
		}
	}

	return false
}
