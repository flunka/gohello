package das

import "testing"

func TestLinearSearch(t *testing.T) {
	cases := []struct {
		in   []int
		x    int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{1, 2, 3, 5}, 6, false},
	}
	for _, c := range cases {
		got := LinearSearch(c.in, c.x)
		if got != c.want {
			t.Errorf("Got %t, want %t, in: LinearSearch(%v, %v)", got, c.want, c.in, c.x)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		in   []int
		x    int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 1, true},
		{[]int{1, 2, 3, 5}, 6, false},
	}
	for _, c := range cases {
		got := BinarySearch(c.in, c.x)
		if got != c.want {
			t.Errorf("Got %t, want %t, in: BinarySearch(%v, %v)", got, c.want, c.in, c.x)
		}
	}
}

func TestCrystalBallSearch(t *testing.T) {
	cases := []struct {
		in   []bool
		want int
	}{
		{[]bool{false, false, false, true, true, true}, 3},
		{[]bool{false, false, false, false, false, true}, 5},
		{[]bool{true, true, true, true, true, true}, 0},
		{[]bool{false, false, false, false, false, false}, -1},
	}
	for _, c := range cases {
		got := CrystalBallSearch(c.in)
		if got != c.want {
			t.Errorf("Got %d, want %d, in: CrystalBallSearch(%v)", got, c.want, c.in)
		}
	}
}
