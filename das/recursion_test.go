package das

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Run("simple maze", func(t *testing.T) {
		maze := []string{
			"######## #",
			"#        #",
			"# ########",
		}
		end := Point{8, 0}
		start := Point{1, 2}
		const wall byte = '#'
		got := Solve(maze, wall, start, end)
		want := []Point{start, {1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}, {6, 1}, {7, 1}, {8, 1}, end}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("complex maze", func(t *testing.T) {
		maze := []string{
			"########E#",
			"###      #",
			"#   ## ###",
			"### ## ###",
			"#   ##   #",
			"#S########",
		}
		end := Point{8, 0}
		start := Point{1, 5}
		const wall byte = '#'
		got := Solve(maze, wall, start, end)
		want := []Point{start, {1, 4}, {2, 4}, {3, 4}, {3, 3}, {3, 2}, {3, 1}, {4, 1}, {5, 1}, {6, 1}, {7, 1}, {8, 1}, end}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})

}
