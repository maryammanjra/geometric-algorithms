package planesweep

import (
	"geometric-algorithms/avl-tree"
	"geometric-algorithms/geometry"
)

type Event struct {
	endpoint geometry.Point
	segment  geometry.Segment
	side     int // 0 for left endpoint, 1 for intersection, 2 for right
}

type Sweep struct {
	X      float64
	Status *avl.AVLTree[geometry.Segment]
}

func NewSweep() *Sweep {
	sweep := &Sweep{}

	cmp := func(a, b geometry.Segment) int {
		y1 := yAtX(a, sweep.X)
		y2 := yAtX(b, sweep.X)

		if y1 < y2 {
			return -1
		}
		if y1 > y2 {
			return 1
		}

		if a.ID < b.ID {
			return -1
		}
		if a.ID > b.ID {
			return 1
		}
		return 0
	}

	sweep.Status = avl.NewAVL[geometry.Segment](cmp)
	return sweep
}

func compareEvents(e1 Event, e2 Event) int {

	if e1.endpoint.X > e2.endpoint.X {
		return 1
	}
	if e1.endpoint.X < e2.endpoint.X {
		return -1
	}

	if e1.endpoint.Y > e2.endpoint.Y {
		return 1
	}
	if e1.endpoint.Y < e2.endpoint.Y {
		return -1
	}

	if e1.side < e2.side {
		return 1
	}
	if e1.side > e2.side {
		return -1
	}

	return 0
}

func yAtX(s geometry.Segment, x float64) float64 {

	dx := s.RightEndpoint.X - s.LeftEndpoint.X

	if dx == 0 {
		return s.LeftEndpoint.Y
	}

	t := (x - s.LeftEndpoint.X) / dx
	return s.LeftEndpoint.Y + t*(s.RightEndpoint.Y-s.LeftEndpoint.Y)
}

func checkForIntersection(s1 geometry.Segment, s2 geometry.Segment) bool {
	return false
}

func findIntersection(s1 geometry.Segment, s2 geometry.Segment) geometry.Point {
	return geometry.Point{X: 1, Y: 1}
}

func planeSweep(lines []geometry.Segment) {

	segmentTree := avl.NewAVL(compareEvents)
	statusTree := NewSweep()

	for _, segment := range lines {
		leftEndpointEvent := Event{endpoint: segment.LeftEndpoint, segment: segment, side: 0}
		rightEndpointEvent := Event{endpoint: segment.RightEndpoint, segment: segment, side: 2}
		segmentTree.Root = avl.Insert(segmentTree.Root, segmentTree.Comparator, leftEndpointEvent)
		segmentTree.Root = avl.Insert(segmentTree.Root, segmentTree.Comparator, rightEndpointEvent)
	}

	statusTree.X = segmentTree.FindSmallest().Value.endpoint.X
	segmentTree.PrintGraphical()

	for !segmentTree.IsEmpty() {
		currMin := segmentTree.FindSmallest().Value
		segmentTree.Root = avl.Delete(segmentTree.Root, segmentTree.Comparator, currMin)

		if currMin.side == 0 {
			statusTree.Status.Root = avl.Insert(statusTree.Status.Root, statusTree.Status.Comparator, currMin.segment)
		}

		segmentTree.PrintGraphical()
	}

}
