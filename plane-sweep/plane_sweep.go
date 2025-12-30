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

func compareEvents(e1 Event, e2 Event) int {

	if e1.endpoint.X > e2.endpoint.X {
		return 1
	}
	if e1.endpoint.X < e2.endpoint.X {
		return -1
	}

	if e1.endpoint.Y < e2.endpoint.Y {
		return 1
	}
	if e1.endpoint.Y > e2.endpoint.Y {
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

/*
For the purpose of comparing segments along the sweep-line for above and below other segments i.e
one segment is above another segment -- consider adding the cases for when the line ends
may or may not be needed based on how events are processed
*/
func compareSegments(s1 geometry.Segment, s2 geometry.Segment) int {

	if s1.LeftEndpoint.Y > s2.LeftEndpoint.Y {
		return 1
	}

	if s1.LeftEndpoint.Y < s2.LeftEndpoint.Y {
		return -1
	}

	if s1.RightEndpoint.Y > s2.RightEndpoint.Y {
		return 1
	}

	if s1.RightEndpoint.Y < s2.RightEndpoint.Y {
		return -1
	}

	return 0
}

func checkForIntersection(s1 geometry.Segment, s2 geometry.Segment) bool {
	return false
}

func findIntersection(s1 geometry.Segment, s2 geometry.Segment) geometry.Point {
	return geometry.Point{X: 1, Y: 1}
}

func planeSweep(lines []geometry.Segment) {

	segmentTree := avl.NewAVL(compareEvents)
	statusTree := avl.NewAVL(compareSegments)

	for _, segment := range lines {
		leftEndpointEvent := Event{endpoint: segment.LeftEndpoint, segment: segment, side: 0}
		rightEndpointEvent := Event{endpoint: segment.RightEndpoint, segment: segment, side: 2}
		segmentTree.Root = avl.Insert(segmentTree.Root, segmentTree.Comparator, leftEndpointEvent)
		segmentTree.Root = avl.Insert(segmentTree.Root, segmentTree.Comparator, rightEndpointEvent)
	}

	segmentTree.PrintGraphical()

	for !segmentTree.IsEmpty() {
		currMin := segmentTree.FindSmallest().Value
		segmentTree.Root = avl.Delete(segmentTree.Root, segmentTree.Comparator, currMin)

		if currMin.side == 0 {
			statusTree.Root = avl.Insert(statusTree.Root, statusTree.Comparator, currMin.segment)
		}

		segmentTree.PrintGraphical()
		statusTree.PrintGraphical()
	}

}
