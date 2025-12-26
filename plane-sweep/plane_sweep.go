package planesweep

import (
	"fmt"
	"geometric-algorithms/avl-tree"
	"geometric-algorithms/geometry"
)

type Event struct {
	endpoint geometry.Point
	segment  geometry.Segment
	side     int // 0 for upper endpoint, 1 for lower endpoint
}

func compareEvents(e1 Event, e2 Event) int {

	if e1.endpoint.Y > e2.endpoint.Y {
		return 1
	}
	if e1.endpoint.Y < e2.endpoint.Y {
		return -1
	}

	if e1.side != e2.side {
		if e1.side == 0 {
			return 1
		}
		return -1
	}

	if e1.endpoint.X < e2.endpoint.X {
		return 1
	}
	if e1.endpoint.X > e2.endpoint.X {
		return -1
	}

	return 0
}

func planeSweep(lines []geometry.Segment) {

	segmentTree := avl.NewAVL(compareEvents)
	for _, segment := range lines {
		upperEndpointEvent := Event{endpoint: segment.UpperEndpoint, segment: segment, side: 0}
		lowerEndpointEvent := Event{endpoint: segment.LowerEndpoint, segment: segment, side: 1}
		segmentTree.Root = avl.Insert(segmentTree.Root, segmentTree.Comparator, upperEndpointEvent)
		segmentTree.Root = avl.Insert(segmentTree.Root, segmentTree.Comparator, lowerEndpointEvent)
	}

	segmentTree.PrintGraphical()

	for !segmentTree.IsEmpty() {
		fmt.Println(segmentTree.FindLargest().Value)
		segmentTree.Root = avl.Delete(segmentTree.Root, segmentTree.Comparator, segmentTree.FindLargest().Value)
		segmentTree.PrintGraphical()
	}

}
