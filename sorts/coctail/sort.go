package coctail

//Asc .. set ascending order for sort func
func Asc(item1, item2 int) bool {
	if item1 > item2 {
		return true
	}
	return false
}

//Desc .. set descending order for sort func
func Desc(item1, item2 int) bool {
	if item1 < item2 {
		return true
	}
	return false
}

//Sort .. return ordered slice of int
func Sort(arr []int, compare func(int, int) bool) []int {
	leftIndex, rightIndex := 0, len(arr)-1
	isShiftingFromLtoR, isShiftingFromRtoL := true, true
	for isShiftingFromLtoR && isShiftingFromRtoL {
		isShiftingFromLtoR, isShiftingFromRtoL = false, false
		for i := leftIndex; i < rightIndex; i++ {
			if compare(arr[i], arr[i+1]) {
				arr[i], arr[i+1], isShiftingFromLtoR = arr[i+1], arr[i], true
			}
		}
		if !isShiftingFromLtoR {
			break
		}
		rightIndex--
		for i := rightIndex; i > leftIndex; i-- {
			if compare(arr[i-1], arr[i]) {
				arr[i], arr[i-1], isShiftingFromRtoL = arr[i-1], arr[i], true
			}
		}
		leftIndex++
	}
	return arr
}
