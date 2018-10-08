package insertion

// Sort returns sorted slice of integers.
// At each iteration, InsertionSort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there. It repeats until no input elements remain.
// See: https://ru.wikipedia.org/wiki/Сортировка_вставками
func Sort(a []int, greaterThen func(int, int) bool) []int {
	var n = len(a)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if greaterThen(a[j-1], a[j]) {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
	return a
}

// Acc is comparision function to sort in accending order
func Acc(a, b int) bool {
	return a > b
}

// Dec is comparision function to sort in accending order
func Dec(a, b int) bool {
	return a < b
}
