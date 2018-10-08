package goMisha

//Func Min:
//func to identifier index of the smallest element in array
func Min(a []int) (k int) {
	mum := a[0]
	for i := range a {
		if a[i] < mum {
			mum = a[i]
			k = i
		}
	}
	return
}

//Func "Max":
//func to identifier index of the biggest element in array.
func Max(a []int) int {
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

//Func:SelectSort
//SelectSort with modifing operation
func SelectSortAppModify(a []int, operation func(k []int) int) []int {
	k := len(a)
	s := make([]int, 0)
	loop := true
	for loop {
		for i := 0; i < k; i++ {
			s = append(s, a[operation(a)])
			a = append(a[:operation(a)], a[operation(a)+1:]...)
		}
		loop = false
	}
	a = append(a, s[:]...)
	return a
}

//additional funcs to compare integers
func more(a, b int) (k bool) {
	if a < b {
		k = true
	}
	return k
}

//additional funcs to compare integers
func less(a, b int) (k bool) {
	if a > b {
		k = true
	}
	return k
}

//Func:SelectSortOperate
//SelectSort for array of integres on the place
func SelectSortOperate(a []int, operation func(a, b int) bool) []int {
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

func goMisha() {

}
