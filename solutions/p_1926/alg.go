package main

import "fmt"

func nearestExit(maze [][]byte, entrance []int) int {
	start := NewPoint(entrance[1], entrance[0], 0)
	closed := map[int]bool{start.GetID(): true}
	opened := make([]*Point, 0, 64)
	opened = append(opened, getLimitNeighbor(maze, start)...)
	var cur *Point
	for len(opened) > 0 {
		cur, opened = dequeue(opened)
		// skip duplicate
		if closed[cur.GetID()] {
			continue
		}
		if isExit(maze, cur) {
			return cur.Distance
		}
		opened = append(opened, getLimitNeighbor(maze, cur)...)
		closed[cur.GetID()] = true
	}
	return -1
}

func getLimitNeighbor(maze [][]byte, p *Point) []*Point {
	left, right, up, down := p.Left(), p.Right(), p.Up(), p.Down()
	neighbors := make([]*Point, 0, 4)
	if inBound(maze, left) && !isWall(maze, left) && !p.Equal(left) {
		neighbors = append(neighbors, left)
	}
	if inBound(maze, right) && !isWall(maze, right) && !p.Equal(left) {
		neighbors = append(neighbors, right)
	}
	if inBound(maze, up) && !isWall(maze, up) && !p.Equal(left) {
		neighbors = append(neighbors, up)
	}
	if inBound(maze, down) && !isWall(maze, down) && !p.Equal(left) {
		neighbors = append(neighbors, down)
	}
	return neighbors
}

func inBound(maze [][]byte, p *Point) bool {
	return p.Y >= 0 && p.Y < len(maze) && p.X >= 0 && p.X < len(maze[p.Y])
}

func isExit(maze [][]byte, p *Point) bool {
	if p.X == 0 || p.X == len(maze[p.Y])-1 {
		return true
	}
	if p.Y == 0 || p.Y == len(maze)-1 {
		return true
	}
	return false
}

func isWall(maze [][]byte, p *Point) bool {
	return maze[p.Y][p.X] == WALL
}

func dequeue(queue []*Point) (*Point, []*Point) {
	item := queue[0]
	return item, queue[1:]
}

const (
	WALL = '+'
)

type Point struct {
	X        int
	Y        int
	Id       int
	Distance int
}

func NewPoint(x, y, dist int) *Point {
	return &Point{
		X:        x,
		Y:        y,
		Id:       -1,
		Distance: dist,
	}
}

func (p *Point) GetID() int {
	if p.Id == -1 {
		p.Id = p.X*100 + p.Y
	}
	return p.Id
}

func (p *Point) Left() *Point {
	return NewPoint(p.X-1, p.Y, p.Distance+1)
}

func (p *Point) Right() *Point {
	return NewPoint(p.X+1, p.Y, p.Distance+1)
}

func (p *Point) Up() *Point {
	return NewPoint(p.X, p.Y-1, p.Distance+1)
}

func (p *Point) Down() *Point {
	return NewPoint(p.X, p.Y+1, p.Distance+1)
}

func (p *Point) Equal(o *Point) bool {
	return p.X == o.X && p.Y == o.Y
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v, %v)(%v)", p.X, p.Y, p.Distance)
}
