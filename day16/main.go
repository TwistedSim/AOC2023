package main

import (
	"aoc/util"
	"flag"
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

type Vector struct {
	x int
	y int
}

func (v1 *Vector) add(v2 Vector) Vector {
	return Vector{v1.x + v2.x, v1.y + v2.y}
}

type Path struct {
	Position  Vector
	Direction Vector
}

func isInBound(v Vector, maxX, maxY int) bool {
	return v.x >= 0 && v.x < maxX && v.y >= 0 && v.y < maxY
}

func propagate(contraption [][]rune, position Vector, direction Vector) int {
	propagation := map[Vector][]Vector{}
	paths := []Path{
		{
			Position:  position,
			Direction: direction,
		},
	}
	for len(paths) > 0 {
		var path Path
		path, paths = paths[0], paths[1:]
		position, direction := path.Position, path.Direction
		for isInBound(position, len(contraption), len(contraption[0])) && !slices.Contains(propagation[position], direction) {
			propagation[position] = append(propagation[position], direction)
			switch contraption[position.x][position.y] {
			case '|':
				if direction.y != 0 {
					d := Vector{-1, 0}
					paths = append(paths, Path{Position: position.add(d), Direction: d})
					direction = Vector{1, 0}
				}
			case '-':
				if direction.x != 0 {
					d := Vector{0, -1}
					paths = append(paths, Path{Position: position.add(d), Direction: d})
					direction = Vector{0, 1}
				}
			case '/':
				direction = Vector{-direction.y, -direction.x}
			case '\\':
				direction = Vector{direction.y, direction.x}
			}
			position = position.add(direction)
		}

	}
	return len(propagation)
}

func part1(input string) (result int) {
	c := util.ParseRuneGrid(input)
	return propagate(c, Vector{0, 0}, Vector{0, 1})
}

func part2(input string) int {
	c := util.ParseRuneGrid(input)
	values := []int{}
	for x := 0; x < len(c); x++ {
		values = append(values, propagate(c, Vector{x, 0}, Vector{0, 1}))
		values = append(values, propagate(c, Vector{x, len(c)-1}, Vector{0, -1}))
	}
	for y := 0; y < len(c[0]); y++ {
		values = append(values, propagate(c, Vector{0, y}, Vector{1, 0}))
		values = append(values, propagate(c, Vector{len(c[0])-1, y}, Vector{-1, 0}))
	}
	_, result := util.MinMax(values)
	return result
}
