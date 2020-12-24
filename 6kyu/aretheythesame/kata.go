package aretheythesame

func Comp(array1 []int, array2 []int) bool {

	if array1 == nil && array2 == nil {
		return true
	}

	if array1 == nil || array2 == nil {
		return false
	}

	for _, a := range array1 {
		square := a * a
		found := false

		for _, b := range array2 {
			if b == square {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}
