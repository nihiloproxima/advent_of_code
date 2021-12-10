package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type description struct {
	startX int
	startY int
	endX   int
	endY   int
}

type point struct {
	x int
	y int
}

func toInteger(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return intValue
}

func getLineDirection(desc description) string {
	if desc.startX != desc.endX && desc.startY != desc.endY {
		return "diagonal"
	} else if desc.startX != desc.endX {
		return "horizontal"
	} else {
		return "vertical"
	}

}

func getPoints(line string, withDiagonals bool) []point {
	segments := strings.Split(line, " -> ")

	var desc description = description{
		startX: toInteger(strings.Split(segments[0], ",")[0]),
		startY: toInteger(strings.Split(segments[0], ",")[1]),
		endX:   toInteger(strings.Split(segments[1], ",")[0]),
		endY:   toInteger(strings.Split(segments[1], ",")[1]),
	}

	lineDirection := getLineDirection(desc)

	xDirection := "forward"
	if desc.startX > desc.endX {
		xDirection = "backward"
	}

	var points []point
	if lineDirection == "horizontal" {
		if xDirection == "forward" {
			for x := desc.startX; x <= desc.endX; x++ {
				points = append(points, point{x: x, y: desc.startY})
			}
		} else {
			for x := desc.startX; x >= desc.endX; x-- {
				points = append(points, point{x: x, y: desc.startY})
			}
		}
	} else if lineDirection == "vertical" {
		if desc.startY < desc.endY {
			for y := desc.startY; y <= desc.endY; y++ {
				points = append(points, point{x: desc.startX, y: y})
			}
		} else {
			for y := desc.startY; y >= desc.endY; y-- {
				points = append(points, point{x: desc.startX, y: y})
			}
		}
	} else if withDiagonals == true {
		yDirection := "down"
		if desc.startY > desc.endY {
			yDirection = "up"
		}

		if yDirection == "down" && xDirection == "forward" {
			x := desc.startX
			for y := desc.startY; y <= desc.endY; y++ {
				points = append(points, point{x, y})
				x++
			}
		} else if yDirection == "down" && xDirection == "backward" {
			x := desc.startX
			for y := desc.startY; y <= desc.endY; y++ {
				points = append(points, point{x, y})
				x--
			}
		} else if yDirection == "up" && xDirection == "forward" {
			x := desc.startX
			for y := desc.startY; y >= desc.endY; y-- {
				points = append(points, point{x, y})
				x++
			}
		} else if yDirection == "up" && xDirection == "backward" {
			x := desc.startX
			for y := desc.startY; y >= desc.endY; y-- {
				points = append(points, point{x, y})
				x--
			}
		}
	}

	return points
}

func countOverlaps(grid [][]int) int {
	sum := 0
	for _, row := range grid {
		for _, value := range row {
			if value >= 2 {
				sum++
			}
		}
	}

	return sum
}

func resolve(lines []string, withDiagonals bool) int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, line := range lines {
		points := getPoints(line, withDiagonals)
		for _, point := range points {
			gridPos := &grid[point.y][point.x]
			*gridPos++
		}
	}

	return countOverlaps(grid)
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(input), "\n")

	fmt.Printf("Part 1: %d\n", resolve(lines, false))
	fmt.Printf("Part 2: %d\n", resolve(lines, true))
}
