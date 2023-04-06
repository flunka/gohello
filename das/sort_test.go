package das

import "testing"

func isSorted(t testing.TB, array []int) {
	t.Helper()
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			t.Errorf("Array %v is not sorted", array)
			return
		}
	}
}

func TestBubbleSort(t *testing.T) {
	case1 := []int{3, 4, 6, 1}
	BubbleSort(case1)
	isSorted(t, case1)
}
