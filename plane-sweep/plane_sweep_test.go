package planesweep

import (
	"geometric-algorithms/geometry"
	"testing"
)

func TestPlaneSweep_Basic(t *testing.T) {
	lines := []geometry.Segment{
		{
			LeftEndpoint:  geometry.Point{X: 1, Y: 5},
			RightEndpoint: geometry.Point{X: 1, Y: 1},
		},
		{
			LeftEndpoint:  geometry.Point{X: 3, Y: 6},
			RightEndpoint: geometry.Point{X: 3, Y: 2},
		},
		{
			LeftEndpoint:  geometry.Point{X: 2, Y: 4},
			RightEndpoint: geometry.Point{X: 5, Y: 0},
		},
	}

	planeSweep(lines)
}
