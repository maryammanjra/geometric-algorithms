package geometry

type Segment struct {
	LeftEndpoint  Point
	RightEndpoint Point
}

func compare(sOne Segment, sTwo Segment) int {

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
