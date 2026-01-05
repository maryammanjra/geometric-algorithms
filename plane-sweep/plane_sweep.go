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

func orientation(p, q, r geometry.Point) int {
	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)

	if val > 0 {
		return 1
	}
	return 2
}

func checkForIntersection(s1, s2 geometry.Segment) bool {
	p1 := s1.LeftEndpoint
	q1 := s1.RightEndpoint
	p2 := s2.LeftEndpoint
	q2 := s2.RightEndpoint

	return orientation(p1, q1, p2) != orientation(p1, q1, q2) &&
		orientation(p2, q2, p1) != orientation(p2, q2, q1)
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

		// First endpoint of a segment, check if it will eventually intersect its neighbours, if so insert into the 
		//event queue 
		if currMin.side == 0 {
			statusTree.Status.Root = avl.Insert(statusTree.Status.Root, statusTree.Status.Comparator, currMin.segment)
			aboveSegment := statusTree.Status.FindLargerNeighbour(currMin.segment).Value
			belowSegment := statusTree.Status.FindSmallerNeighbour(currMin.segment).Value

			if checkForIntersection(aboveSegment, currMin.segment){
				intersectionPointAbove := findIntersection(aboveSegment, currMin.segment)
				intersectionAboveEvent := Event{endpoint: intersectionPointAbove, segment: currMin.segment, side: 1}
				segmentTree.Root = avl.Insert(segmentTree.Root, segmentTree.Comparator, intersectionAboveEvent)
			}

			if checkForIntersection(belowSegment, currMin.segment) {
				intersectionPointBelow := findIntersection(aboveSegment, currMin.segment)
				intersectionBelowEvent := Event{endpoint: intersectionPointBelow, segment: currMin.segment, side: 1}
				segmentTree.Root = avl.Insert(segmentTree.Root, segmentTree.Comparator, intersectionBelowEvent)
			}	
		}

		// Intersection point of a segment, check which segments it intersects out of its neighbours, then change the ordering
		// in the status tree 
		else if currMin.side == 1 {
			
		}

		segmentTree.PrintGraphical()
	}

}
