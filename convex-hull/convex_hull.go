package convexhull

import (
	"geometric-algorithms/geometry"
	"math"
	"sort"
)

// Move Stack and Point delcarations
type Stack struct {
	stack []geometry.Point
}

func (s *Stack) push(p geometry.Point) {
	s.stack = append(s.stack, p)
}

func (s *Stack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) pop() geometry.Point {
	if !s.isEmpty() {
		popped := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return popped
	}
	return geometry.Point{X: -1, Y: -1}
}

func (s *Stack) peek() geometry.Point {
	if !s.isEmpty() {
		last := s.stack[len(s.stack)-1]
		return last
	}
	return geometry.Point{X: -1, Y: -1}
}

func orientedArea(p1 geometry.Point, p2 geometry.Point, p3 geometry.Point) float64 {
	vectorOne := geometry.Point{X: p2.X - p1.X, Y: p2.Y - p1.Y}
	vectorTwo := geometry.Point{X: p3.X - p1.X, Y: p3.Y - p1.Y}
	return (vectorOne.X * vectorTwo.Y) - (vectorOne.Y * vectorTwo.X)
}

// findPolarAngle assumes p2 lies either in quadrant one or quadrant two relative to p1 as origin,
// based on Graham's scan calculating polar angles in reference to the lowest Y-coordinate.
func findPolarAngle(p1 geometry.Point, p2 geometry.Point) float64 {
	xVector := p2.X - p1.X
	yVector := p2.Y - p1.Y

	if xVector < 0 {
		return math.Pi + math.Atan(yVector/xVector)
	}

	return math.Atan(yVector / xVector)
}

func findSmallestY(points []geometry.Point) geometry.Point {
	minPoint := points[0]

	for _, val := range points {
		if val.Y < minPoint.Y {
			minPoint = val
		}
	}

	return minPoint
}

func sortByPolarAngle(points []geometry.Point, smallestY geometry.Point) {
	sort.Slice(points, func(i, j int) bool {
		return findPolarAngle(smallestY, points[i]) < findPolarAngle(smallestY, points[j])
	})
}

func grahamsScan(points []geometry.Point) []geometry.Point {
	startingPoint := findSmallestY(points)
	sortByPolarAngle(points, startingPoint)
	stack := new(Stack)

	pointOne := points[0]
	pointTwo := points[1]

	stack.push(pointOne)
	stack.push(pointTwo)

	for i := 2; i < len(points); i++ {
		if orientedArea(pointOne, pointTwo, points[i]) > 0 {
			stack.push(points[i])
			pointOne = pointTwo
			pointTwo = points[i]
		} else if orientedArea(pointOne, pointTwo, points[i]) < 0 {
			for orientedArea(pointOne, pointTwo, points[i]) < 0 {
				stack.pop()
				pointTwo = pointOne
				pointOne = stack.stack[len(stack.stack)-2]
			}
			stack.push(points[i])
		}
	}

	return stack.stack
}
