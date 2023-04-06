package das

type Point struct {
	x, y int
}

func Solve(maze []string, wall byte, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}
	var path = make([]Point, 0, 10)
	walk(maze, wall, start, end, seen, &path)
	return path
}

var dir = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func walk(maze []string, wall byte, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, curr)
		return true
	}
	if curr.y >= len(maze) || curr.y < 0 || curr.x >= len(maze[0]) || curr.x < 0 {
		return false
	}
	if maze[curr.y][curr.x] == wall {
		return false
	}
	if seen[curr.y][curr.x] {
		return false
	}

	seen[curr.y][curr.x] = true
	*path = append(*path, curr)
	for i := 0; i < len(dir); i++ {
		step := dir[i]
		if walk(maze, wall, Point{curr.x + step[0], curr.y + step[1]}, end, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}
