package main 

var Dir = []Point{{1,0}, {0,1}, {-1, 0}, {0, -1}}

type Point struct {
    x int
    y int
}
 

func main() {
    maze := []string{
        "xxxxxxxxxx x",
        "x        x x",
        "x        x x",
        "x xxxxxxxx x",
        "x          x",
        "x xxxxxxxxxx",
    }
    start := Point{5,1}
    final := Point{0,10}

    result := solve(maze, start, final)

    for _, b := range result {
        println(b.y, b.x)
    }
}

func solve(maze []string, start Point, final Point) []Point {
    positions := make([]Point, 0)
    been := &[6][12]bool{}
    
    walk(maze, start, final, &positions, been)
    return positions
}

func walk(maze []string, curr Point, end Point, positions *[]Point, been *[6][12]bool) bool {
    if curr.x < 0 || curr.x >= len(maze) || curr.y >= len(maze[0]) || curr.y < 0  {
        return false
    }

    if maze[curr.x][curr.y] == 'x' {
        return false
    }

    if curr.x == end.x && curr.y == end.y {
        *positions = append(*positions, curr)
        return true
    }

    if been[curr.x][curr.y] {
        return false
    }
    
    been[curr.x][curr.y] = true

    for _, dir := range Dir {
        if walk(maze, Point{curr.x + dir.x, curr.y + dir.y}, end, positions, been) {
            *positions = append(*positions, curr)
            return true
        }
    }

    return false
}
