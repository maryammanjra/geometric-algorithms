package planesweep

import (
	"geometric-algorithms/avl-tree"
	"geometric-algorithms/geometry"
)

func compareSegments(sOne geometry.Segment, sTwo geometry.Segment) int {

	if (sOne.LeftEndpoint.Y > sTwo.LeftEndpoint.Y) && (sOne.LeftEndpoint.Y > sTwo.RightEndpoint.Y) ||
		(sOne.RightEndpoint.Y > sTwo.LeftEndpoint.Y) && (sOne.RightEndpoint.Y > sTwo.RightEndpoint.Y) {
		return 1
	} else if (sTwo.LeftEndpoint.Y > sOne.LeftEndpoint.Y) && (sTwo.LeftEndpoint.Y > sOne.RightEndpoint.Y) ||
		(sTwo.RightEndpoint.Y > sOne.LeftEndpoint.Y) && (sTwo.RightEndpoint.Y > sOne.RightEndpoint.Y) {
		return -1
	} else {
		return 0
	}

}

func planeSweep(lines []geometry.Segment) {
	segmentTree := avl.NewAVL(compareSegments)
	for _, segment := range lines {
		segmentTree.Root = avl.Insert(segmentTree.Root, segmentTree.Comparator, segment)
	}
}
