package planesweep

import (
	"geometric-algorithms/geometry"
	"testing"
)

func TestPlaneSweep_Basic(t *testing.T) {
	lines := []geometry.Segment{
		{
			UpperEndpoint: geometry.Point{X: 1, Y: 5},
			LowerEndpoint: geometry.Point{X: 1, Y: 1},
		},
		{
			UpperEndpoint: geometry.Point{X: 3, Y: 6},
			LowerEndpoint: geometry.Point{X: 3, Y: 2},
		},
		{
			UpperEndpoint: geometry.Point{X: 2, Y: 4},
			LowerEndpoint: geometry.Point{X: 5, Y: 0},
		},
	}

	planeSweep(lines)
}
