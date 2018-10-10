package selection

//asc to identifier index of the smallest element in array
func asc(a []int) int {
	if len(a) == 0 {
		return 0
	}
	mum := a[0]
	var k int
	for i := range a {
		if a[i] < mum {
			mum = a[i]
			k = i

		}
	}
	return k
}

//desc to identifier index of the biggest element in array.
func desc(a []int) int {
	if len(a) == 0 {
		return 0
	}
	mum := a[0]
	var k int
	for i := range a {
		if a[i] > mum {
			mum = a[i]
			k = i
		}
	}

	return k
}

//Sort array with additional slice
func Sort(a []int, operation func(k []int) int) []int {

	for {
		k := len(a)
		s := []int{}
		for i := 0; i < k; i++ {
			s = append(s, a[operation(a)])
			a = append(a[:operation(a)], a[operation(a)+1:]...)
		}
		a = append(a, s[:]...)
		break
	}

	return a
}

//bigger funcs to compare integers
func bigger(a, b int) (k bool) {
	if a > b {
		k = true

	}
	return k
}

//less funcs to compare integers
func less(a, b int) (k bool) {
	if a < b {
		k = true
	}
	return k
}

//SortOnPlace sort array of int
func SortOnPlace(a []int, operation func(a, b int) bool) []int {
	k := len(a)
	for i := 0; i < k; i++ {
		indMin := i
		for j := i; j < k; j++ {
			if operation(a[j], a[indMin]) {
				indMin = j
			}
		}
		a[i], a[indMin] = a[indMin], a[i]
	}
	return a
}

func selection() {

}
